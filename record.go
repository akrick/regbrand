package main

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
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
	uline, err = r.GbkToUtf8([]byte(line))
	if err != nil{
		return
	}
	return
}
func  (r *Record) GbkToUtf8(str []byte) (rs []byte, err error) {
	rd := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder())
	rs, err = ioutil.ReadAll(rd)
	if err != nil {
		return
	}
	return
}