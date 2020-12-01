package main

import (
	"fmt"
	"github.com/gocolly/colly"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type QdsScraper struct {
	Ql interface{}
	Url string
	Cookie string
	No int32
	Catalog int32
}

func (qds *QdsScraper) SetCookie(cookie string) {
	qds.Cookie = cookie
}
func (qds *QdsScraper) SetUrl(url string) {
	qds.Url = url
}
func (qds *QdsScraper) Search(period int, category int, pageNum int) (err error){

	if pageNum > 0 {
		c := colly.NewCollector()
		for i := 1; i <= pageNum; i++{
			url := qds.Url + "?&period="+strconv.Itoa(period)+"&category="+strconv.Itoa(category)+"&page="+strconv.Itoa(i)
			c.OnRequest(func(r *colly.Request) {
				r.Headers.Set("cookie", qds.Cookie)
				r.Ctx.Put("brandName", "")
				r.Ctx.Put("status", "all")
				r.Ctx.Put("applicationId", "")
				r.Ctx.Put("applicationName", "")
				r.Ctx.Put("agency", "")
				r.Ctx.Put("noticeTime", "")
				r.Ctx.Put("typeCode", "")
				r.Ctx.Put("page", i)
			})

			c.OnResponse(func(res *colly.Response) {

				html := res.Body
				rs := jsoniter.Get(html[:], "data").Get("data").Get("items").ToString()
				var json = jsoniter.ConfigCompatibleWithStandardLibrary
				var result []QdsItem
				err = json.Unmarshal([]byte(rs), &result)
				if err != nil {
					log.Fatal(err)
					return
				}
				for _, item := range result {
					var record Record
					item.Link, _ = qds.FetchImage(item.RegNo, item.TmName)
					tycItem, err := qds.FetchContactInfo(item.ApplicantCn)
					if err != nil{
						log.Fatal(err)
						return
					}
					record.ApplicationCn = item.ApplicantCn
					record.RegLocation = tycItem.RegLocation
					record.Link = item.Link
					record.PhoneNumber = tycItem.PhoneNumber
					record.LegalPersonName = tycItem.LegalPersonName
					record.TmName = item.TmName
					record.IntCls = item.IntCls
					record.RegNo = item.RegNo
					record.AnnouncementIssue = item.AnnouncementIssue

					fmt.Println(record)
					err = qds.PutData(record)
					if err != nil {
						log.Fatal(err)
						return
					}
				}
			})

			c.OnError(func(res *colly.Response, err error) {
				log.Fatal(err)
				return
			})
			err = c.Visit(url)
			if err != nil {
				log.Fatal(err)
				return
			}
		}
	}
	return
}

func (qds *QdsScraper) FetchImage(id string, name string)  (link string, err error) {
	url := "https://so.quandashi.com/search/notice/notice-detail"
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("cookie", qds.Cookie)
		r.Ctx.Put("id", id)
		r.Ctx.Put("issue", qds.No)
		r.Ctx.Put("name", name)
	})
	c.OnHTML("body > div.page.pt-search > div.w-center > div.page-detail > div.content > img", func(e *colly.HTMLElement) {
		link = e.Attr("src")
	})
	c.OnError(func(res *colly.Response, err error) {
		log.Fatal(err)
		return
	})
	c.OnResponse(func(res *colly.Response) {

	})
	err = c.Visit(url)
	if err != nil{
		log.Fatal(err)
		return
	}
	return
}

func (qds *QdsScraper) FetchContactInfo(name string) (tycItem TycMessageItem, err error){
	tyc := new(TianYanCha)
	tyc.SetToken("eab5ec28-886d-4079-99f1-f6b80e00f29a")
	info, err := tyc.GetMessageByUrlToken(name)
	if err != nil{
		log.Fatal(err)
		return
	}
	for _, item := range info.Result{
		tycItem = item
		break
	}
	return
}

func (qds *QdsScraper) PutData(data Record) (err error) {
	path := "./data"
	_, err = os.Stat(path)
	if os.IsNotExist(err){
		err = os.Mkdir(path, 0755)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	csvFile := path + "/data.csv"
	var f *os.File
	f, err = os.OpenFile(csvFile, os.O_APPEND|os.O_CREATE, 0755)

	if err != nil {
		log.Fatal(err)
	}
	uline, err := data.ToString()
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = f.Write(uline)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	//download image
	imageBytes, err := qds.FetchImageByUrl(data.Link)
	if err != nil {
		log.Fatal(err)
		return
	}
	//filter
	data.ApplicationCn = strings.Replace(data.ApplicationCn, "/", "", -1)
	imageFile := "./data/"+data.RegNo+data.ApplicationCn+".jpg"
	f, err = os.OpenFile(imageFile, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = f.Write(imageBytes)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err = f.Close(); err != nil {
		log.Fatal(err)
		return
	}
	return
}

func (qds *QdsScraper) FetchImageByUrl(url string) (pix []byte, err error) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	pix, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return
}
