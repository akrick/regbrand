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
		for i := 0; i <= pageNum; i++{
			var urls []string
			query := "y"+strconv.Itoa(period)+"z"+strconv.Itoa(category)+"w"+strconv.Itoa(i)
			url := qmx.Url + query
			c.OnRequest(func(r *colly.Request) {
				//r.Headers.Set("cookie", qmx.Cookie)
			})
			c.OnHTML("#__layout > div > div.pageIndex.wap > section > div.list-anni.wap > ul > li > div.m-box > div", func(e *colly.HTMLElement) {
				urls = e.ChildAttrs("a", "href")
				if urls != nil {
					for _, item := range urls{
						link := "https://sbgg.qmxip.com"+item
						c.OnHTML("#__layout > div > div.list > div > div.notice-content > div.fl.lbox > div.tm-header.parts > div.tm_n_r.box.fl > a > b", func(el *colly.HTMLElement) {
							//regName := el.Text
						})
						c.OnHTML("#__layout > div > div.list > div > div.notice-content > div.fl.lbox > div:nth-child(2) > div.notice-pattern.mb20 > div.pattern-img.pc > div.fr > ul > li > img", func(el *colly.HTMLElement) {
							//regImage := el.Attr("src")
						})
						c.Visit(link)
					}
				}
			})

			c.OnResponse(func(res *colly.Response) {

				html := res.Body
				rs := jsoniter.Get(html[:], "data").Get("data").Get("items").ToString()
				var json = jsoniter.ConfigCompatibleWithStandardLibrary
				var result []QdsItem
				err = json.Unmarshal([]byte(rs), &result)
				if err != nil{
					log.Fatal(err)
				}
				for _, item := range result {
					var (
						record Record
						tycMsg TycMessage
					)
					item.Link, _ = qmx.FetchImage(item.RegNo, period, item.TmName)

					hkey := qmx.GenerateHashKey(item.ApplicantCn)
					hjson, err := rdb.Get(context.Background(), hkey).Result()
					if err == nil  || err == redis.Nil{//if hkey not exist
						tycMsg, err = qmx.FetchContactInfo(item.ApplicantCn)
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
						record.ApplicationCn = item.ApplicantCn
						record.RegLocation = tycItem.RegLocation
						record.Link = item.Link
						record.PhoneNumber = tycItem.PhoneNumber
						record.LegalPersonName = tycItem.LegalPersonName
						record.TmName = item.TmName
						record.IntCls = item.IntCls
						record.RegNo = item.RegNo
						record.AnnouncementIssue = item.AnnouncementIssue

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
						fmt.Println(item.RegNo+" "+item.ApplicantCn+" "+item.TmName+" "+item.Link)
					}else{
						//do nothing
						//fmt.Println(tycMsg)
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
	path := "./data"
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

