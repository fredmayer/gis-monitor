package models

import "time"

type ContentOfLotDto struct {
	Content    []LotDto
	TotalPages int `json:"totalPages"`
}

type LotDto struct {
	ID           string `json:"id"`
	NoticeNumber string `json:"noticeNumber"`
	LotNumber    int    `json:"lotNumber"`
	LotStatus    string `json:"lotStatus"`
	BiddType     struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"biddType"`
	BiddForm struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"biddForm"`
	LotName        string    `json:"lotName"`
	LotDescription string    `json:"lotDescription"`
	PriceMin       float64   `json:"priceMin"`
	PriceFin       float64   `json:"priceFin"`
	BiddEndTime    time.Time `json:"biddEndTime"`
	LotImages      []string  `json:"lotImages"`
	//Characteristics []struct {
	//	CharacteristicValue string `json:"characteristicValue,omitempty"`
	//	Name                string `json:"name"`
	//	Code                string `json:"code"`
	//	Type                string `json:"type"`
	//	Unit                struct {
	//		Code   string `json:"code"`
	//		Name   string `json:"name"`
	//		Symbol string `json:"symbol"`
	//	} `json:"unit,omitempty"`
	//} `json:"characteristics"`
	CurrencyCode  string `json:"currencyCode"`
	EtpCode       string `json:"etpCode"`
	SubjectRFCode string `json:"subjectRFCode"`
	Category      struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"category"`
	CreateDate     time.Time `json:"createDate"`
	TimeZoneName   string    `json:"timeZoneName"`
	TimezoneOffset string    `json:"timezoneOffset"`
	HasAppeals     bool      `json:"hasAppeals"`
	IsStopped      bool      `json:"isStopped"`
	//Attributes     []struct {
	//	Code          string `json:"code"`
	//	FullName      string `json:"fullName"`
	//	Value         string `json:"value,omitempty"`
	//	AttributeType string `json:"attributeType"`
	//	Group         struct {
	//		Code             string `json:"code"`
	//		Name             string `json:"name"`
	//		DisplayGroupType string `json:"displayGroupType"`
	//	} `json:"group"`
	//	SortOrder int `json:"sortOrder"`
	//} `json:"attributes"`
	//NoticeAttributes []struct {
	//	Code          string `json:"code"`
	//	FullName      string `json:"fullName"`
	//	Value         bool   `json:"value"`
	//	AttributeType string `json:"attributeType"`
	//	Group         struct {
	//		Code             string `json:"code"`
	//		Name             string `json:"name"`
	//		DisplayGroupType string `json:"displayGroupType"`
	//	} `json:"group"`
	//	SortOrder int `json:"sortOrder"`
	//} `json:"noticeAttributes"`
	IsAnnulled                        bool      `json:"isAnnulled"`
	NoticeFirstVersionPublicationDate time.Time `json:"noticeFirstVersionPublicationDate"`
	LotVat                            struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"lotVat"`
	NpaHintCode     string `json:"npaHintCode"`
	TypeTransaction string `json:"typeTransaction"`
}
