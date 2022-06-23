package models

type Numbers struct {
	ID          int    `json:"id"`
	Country     string `json:"country"`
	Valid       string `json:"valid"`
	CountryCode string `json:"country_code"`
	Num         string `json:"num"`
}
