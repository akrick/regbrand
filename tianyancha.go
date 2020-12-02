package main

import (
	"github.com/gocolly/colly"
	jsoniter "github.com/json-iterator/go"
	"log"
	url2 "net/url"
)

type TianYanCha struct {
	Token string
	Name string
}

func (tyc *TianYanCha) SetToken(token string)  {
	tyc.Token = token
}

func (tyc *TianYanCha) GetMessageByUrlToken(name string)  (data TycMessage, err error){

	params := url2.Values{}
	params.Add("name", name)
	params.Add("keyword", name)
	query := params.Encode()
	var result TycMessage
	url := "http://open.api.tianyancha.com//services/open/ic/baseinfoV2/2.0?"+query
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Authorization", tyc.Token)
	})
	c.OnResponse(func(res *colly.Response) {
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
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