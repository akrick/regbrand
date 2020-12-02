package main

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"strconv"
)

type Record struct {
	AnnouncementIssue int
	RegNo string
	TmName string
	IntCls int
	ApplicationCn string
	LegalPersonName string
	PhoneNumber string
	RegLocation string
	Link string
}

func (r *Record) ToString() (uline []byte, err error) {
	line := ""
	line += strconv.Itoa(r.AnnouncementIssue)+","+r.RegNo+","+r.TmName+","+strconv.Itoa(r.IntCls)+","+r.ApplicationCn+","+r.LegalPersonName+","+r.PhoneNumber+","+r.RegLocation+","+r.Link+"\n"
	uline, err = r.Utf8ToGbk([]byte(line))
	if err != nil{
		log.Fatal(err)
		return
	}
	return
}
func (r *Record) GbkToUtf8(str []byte) (rs []byte, err error) {
	rd := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder())
	rs, err = ioutil.ReadAll(rd)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func (r *Record) Utf8ToGbk(str []byte) (rs []byte, err error) {
	rd := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewEncoder())
	rs, err = ioutil.ReadAll(rd)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}