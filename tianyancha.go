package main

import (
	"fmt"
	"github.com/gocolly/colly"
	jsoniter "github.com/json-iterator/go"
	"log"
)

type TianYanCha struct {
	Token string
	Name string
}

func (tyc *TianYanCha) SetToken(token string)  {
	tyc.Token = token
}

func (tyc *TianYanCha) GetMessageByUrlToken(name string)  (data TycMessage, err error){

	var result TycMessage
	url := "http://open.api.tianyancha.com//services/open/ic/baseinfoV2/2.0"
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Authorization", tyc.Token)
		fmt.Println(name)
		r.Ctx.Put("name", name)
		r.Ctx.Put("keyword", name)
	})
	c.OnResponse(func(res *colly.Response) {
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		//fmt.Println(string(res.Body))
		//os.Exit(0)
		err := json.Unmarshal(res.Body, &result)
		if err != nil{
			log.Fatal(err)
		}
	})
	err = c.Visit(url)
	if err != nil{
		log.Fatal(err)
	}
	return result, nil
}