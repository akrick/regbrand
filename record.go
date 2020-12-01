package main

import "strconv"

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

func (r *Record) ToString() string {
	line := ""
	line += strconv.Itoa(r.AnnouncementIssue)+","+r.RegNo+","+r.TmName+","+strconv.Itoa(r.IntCls)+","+r.ApplicationCn+","+r.LegalPersonName+","+r.PhoneNumber+","+r.RegLocation+","+r.Link+"\n"
	return line
}