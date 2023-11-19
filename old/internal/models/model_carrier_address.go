package models

type CarrierAddress struct {
	ID        string `json:"id"`
	CarrierID int64  `json:"carrier_id"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zip_code"`
	Country   string `json:"country"`
}
