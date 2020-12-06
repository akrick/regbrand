package main

type TycMessage struct {
	Reason string `json:"reason"`
	ErrorCode int `json:"error_code"`
	Result *TycMessageItem
}

type TycMessageItem struct {
	PercentileScore int `json:"percentileScore"`
	StaffNumRange string `json:"staffNumRange"`
	FromTime int `json:"fromTime"`
	Type int `json:"type"`
	BondName string `json:"bondName"`
	Id int `json:"id"`
	IsMicroEnt int `json:"isMicroEnt"`
	UsedBondName string `json:"usedBondName"`
	RegNumber string `json:"regNumber"`
	RegCapital string `json:"regCapital"`
	Name string `json:"name"`
	RegInstitute string `json:"regInstitute"`
	RegLocation string `json:"regLocation"`
	ApprovedTime int `json:"approvedTime"`
	UpdateTimes int `json:"updateTimes"`
	SocialStaffNum int `json:"socialStaffNum"`
	Tags string `json:"tags"`
	TaxNumber string `json:"taxNumber"`
	BusinessScope string `json:"businessScope"`
	Property3 string `json:"property3"`
	Alias string `json:"alias"`
	OrgNumber string `json:"orgNumber"`
	RegStatus string `json:"regStatus"`
	EstiblishTime int `json:"estiblishTime"`
	BondType string `json:"bondType"`
	LegalPersonName string `json:"legalPersonName"`
	ToTime int `json:"toTime"`
	ActualCapital string `json:"actualCapital"`
	CompanyOrgType string `json:"companyOrgType"`
	Base string `json:"base"`
	CreditCode string `json:"creditCode"`
	HistoryNames string `json:"historyName"`
	HistoryNameList []string `json:"historyNameList"`
	BondNum string `json:"bondNum"`
	RegCapitalCurrency string `json:"regCapitalCurrency"`
	ActualCapitalCurrency string `json:"actualCapitalCurrency"`
	Email string `json:"email"`
	WebsiteList string `json:"websiteList"`
	PhoneNumber string `json:"phoneNumber"`
	RevokeDate int `json:"revokeDate"`
	RevokeReason string `json:"revokeReason"`
	CancelDate int `json:"cancelDate"`
	CancelReason string `json:"cancelReason"`
	City string `json:"city"`
	District string `json:"district"`
	IndustryAll *Industry
	Reason string `json:"reason"`
	ErrorCode int `json:"error_code"`
}

type Industry struct {
	Category string `json:"category"`
	CategoryBig string `json:"categoryBig"`
	CategoryMiddle string `json:"categoryMiddle"`
	CategorySmall string `json:"categorySmall"`
}

//type HistoryName struct {
//	Name string `json:"name"`
//}