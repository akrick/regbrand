package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gocolly/colly"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"os"
	"strconv"
	"strings"
)

type QmxScraper struct {
	Ql interface{}
	Url string
	Cookie string
	No int
	Catalog int
}

func (qmx *QmxScraper) SetCookie(cookie string) {
	qmx.Cookie = cookie
}
func (qmx *QmxScraper) SetUrl(url string) {
	qmx.Url = url
}
func (qmx *QmxScraper) GenerateHashKey(name string)  (hkey string){
	h := md5.New()
	h.Write([]byte(name))
	hkey = hex.EncodeToString(h.Sum(nil))
	return
}
func (qmx *QmxScraper) Search(period int, category int, pageNum int) (err error){
	qmx.No = period
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // set password
		DB:       0,  // use default DB
	})
	if pageNum > 0 {
		c := colly.NewCollector()
		for i := 1; i <= pageNum; i++{
			var urls []string
			query := "y"+strconv.Itoa(period)+"z"+strconv.Itoa(category)+"w"+strconv.Itoa(i)+"/"
			url := qmx.Url + query
			c.OnRequest(func(r *colly.Request) {
				r.Headers.Set("cookie", qmx.Cookie)
			})
			c.OnHTML("#__layout > div > div.pageIndex.wap > section > div.list-anni.wap > ul", func(e *colly.HTMLElement) {
				e.ForEach("div.m-box > div > a", func(n int, el *colly.HTMLElement) {
					urls = append(urls, el.Request.AbsoluteURL(el.Attr("href")))
				})
				if size := len(urls); size > 0 {
					for _, item := range urls{
						qmxItem := new(QmxItem)
						qmxItem.Period = period
						qmxItem.Category = category
						qmxItem.Link = item
						c.OnRequest(func(r *colly.Request) {
							r.Headers.Set("cookie", qmx.Cookie)
						})
						c.OnHTML("#__layout > div > div.list > div > div.notice-content > div.fl.lbox > div.tm-header.parts > div.tm_n_r.box.fl > a > b", func(el *colly.HTMLElement) {
							qmxItem.Brand = strings.ReplaceAll(strings.ReplaceAll(el.Text, " ", ""), "\n", "")
						})
						c.OnHTML("#__layout > div > div.list > div > div.notice-content > div.fl.lbox > div.tm-header.parts > div.tm_n_r.box.fl > div > span:nth-child(2)", func(el *colly.HTMLElement) {
							qmxItem.RegNo = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(el.Text, " ", ""), "申请/注册号：\n", ""), "\n", "")
						})
						c.OnHTML("#__layout > div > div.list > div > div.notice-content > div.fl.lbox > div:nth-child(2) > div.notice-pattern.mb20 > div.pattern-img.pc > div.fr > ul > li > img", func(el *colly.HTMLElement) {
							qmxItem.Image = el.Attr("src")
						})
						c.OnHTML("#__layout > div > div.list > div > div.notice-content > div.fl.lbox > div:nth-child(2) > div:nth-child(3) > div > p:nth-child(8)", func(el *colly.HTMLElement) {
							qmxItem.Applyer = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(el.Text, " ", ""), "申请人：\n", ""), "\n", "")
						})
						err = c.Visit(qmxItem.Link)
						if err != nil{
							log.Fatal(err)
						}
						//fmt.Println(qmxItem)
						//os.Exit(0)
						var tycMsg TycMessage
						var record Record
						var json = jsoniter.ConfigCompatibleWithStandardLibrary
						tycMsg, err = qmx.FetchContactInfo(qmxItem.Applyer)
						if err != nil{
							log.Fatal(err)
						}

						hkey := qmx.GenerateHashKey(qmxItem.Applyer)
						hjson, err := rdb.Get(context.Background(), hkey).Result()
						if err == nil  || err == redis.Nil{//if hkey not exist
							tycMsg, err = qmx.FetchContactInfo(qmxItem.Applyer)
							if err != nil{
								log.Fatal(err)
							}
						}else{
							err = json.UnmarshalFromString(hjson, tycMsg)
							if err != nil{
								log.Fatal(err)
							}
						}
						//if found contact info
						if tycMsg.ErrorCode == 0 {
							tycItem := tycMsg.Result
							record.ApplicationCn = qmxItem.Applyer
							record.RegLocation = tycItem.RegLocation
							record.Link = qmxItem.Image
							record.PhoneNumber = tycItem.PhoneNumber
							record.LegalPersonName = tycItem.LegalPersonName
							record.TmName = qmxItem.Brand
							record.IntCls = qmxItem.Category
							record.RegNo = qmxItem.RegNo
							record.AnnouncementIssue = qmxItem.Period

							hjson, err := json.MarshalToString(record)
							if err != nil {
								log.Fatal(err)
							}
							err = rdb.Set(context.Background(), hkey, hjson, 0).Err()
							if err != nil {
								log.Fatal(err)
							}
							if strings.Index(record.ApplicationCn, "公司") > 0 {
								err = qmx.PutData(record)
								if err != nil {
									log.Fatal(err)
								}
							}
							fmt.Println(qmxItem.RegNo + " " + qmxItem.Applyer + " " + qmxItem.Brand + " " + qmxItem.Image)
						}
					}
				}
			})

			c.OnError(func(res *colly.Response, err error) {
				log.Fatal(err)
			})
			err = c.Visit(url)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return
}

func (qmx *QmxScraper) FetchImage(id string, period int, name string)  (link string, err error) {
	params := url2.Values{}
	params.Add("id", id)
	params.Add("issue", strconv.Itoa(period))
	params.Add("name", name)
	query := params.Encode()
	url := "https://so.quandashi.com/search/notice/notice-detail?"+query
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("cookie", qmx.Cookie)
	})
	c.OnHTML("body > div.page.pt-search > div.w-center > div.page-detail > div.content > img", func(e *colly.HTMLElement) {
		link = e.Attr("src")
	})
	c.OnError(func(res *colly.Response, err error) {
		log.Fatal(err)
	})
	err = c.Visit(url)
	if err != nil{
		log.Fatal(err)
	}
	return
}

func (qmx *QmxScraper) FetchContactInfo(name string) (tycMsg TycMessage, err error){
	tyc := new(TianYanCha)
	tyc.SetToken("eab5ec28-886d-4079-99f1-f6b80e00f29a")
	tycMsg, err = tyc.GetMessageByUrlToken(name)
	if err != nil{
		log.Fatal(err)
	}
	return
}

func (qmx *QmxScraper) PutData(data Record) (err error) {
	path := "./qmxdata"
	_, err = os.Stat(path)
	if os.IsNotExist(err){
		err = os.Mkdir(path, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	var f *os.File
	//download image
	imageBytes, err := qmx.FetchImageByUrl(data.Link)
	if err != nil {
		log.Fatal(err)
	}
	//filter
	data.ApplicationCn = strings.Replace(data.ApplicationCn, "/", "", -1)
	imageFile := "./data/"+data.RegNo+data.ApplicationCn+".jpg"
	f, err = os.OpenFile(imageFile, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(imageBytes)
	if err != nil {
		log.Fatal(err)
	}
	if err = f.Close(); err != nil {
		log.Fatal(err)
	}

	//write csv record
	csvFile := path + "/data.csv"
	f, err = os.OpenFile(csvFile, os.O_APPEND|os.O_CREATE, 0755)

	if err != nil {
		log.Fatal(err)
	}
	uline, err := data.ToString()
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(uline)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	return
}

func (qmx *QmxScraper) FetchImageByUrl(url string) (pix []byte, err error) {
	if strings.Index(url, "http") < 0{
		url = "https://tm-images.oss-cn-beijing.aliyuncs.com/png/"+url
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(url)
		log.Fatal(err)
	}
	defer resp.Body.Close()
	pix, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return
}

