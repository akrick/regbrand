package main

type QdsItem struct {
	RegNo string `json:"regNo"`
	Year int `json:"year"`
	DetailId string `json:"detailId"`
	RegDate string `json:"regDate"`
	AppDate string `json:"appDate"`
	IntCls int `json:"intCls"`
	ImageUrl string `json:"imageUrl"`
	StatusName string `json:"statusName"`
	StatusZh string `json:"statusZh"`
	AddressEn string `json:"addressEn"`
	Tag int `json:"tag"`
	PrivateEndDate string `json:"privateEndDate"`
	Group int `json:"group"`
	PrivateStartDate string `json:"privateStartDate"`
	AnnouncementIssue int `json:"announcementIssue"`
	Address string `json:"address"`
	Agency string `json:"agency"`
	ApplicantShare string `json:"applicantShare"`
	AnnouncementDate string `json:"announcementDate"`
	ApplicantCn string `json:"applicantCn"`
	TypeFlag int `json:"typeFlag"`
	EnApplicant string `json:"enApplicant"`
	TmName string `json:"tmName"`
	RegIssue int `json:"regIssue"`
	Link string `json:"link"`
}
