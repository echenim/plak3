package models

type BillingInformation struct {
	Id                   int64  `json:"id"`
	ShipperID            int64  `json:"shipper_id"`
	BillingStreet        string `json:"billing_street"`
	BillingCity          string `json:"billing_city"`
	BillingState         string `json:"billing_state"`
	BillingZipCode       string `json:"billing_zip_code"`
	BillingCountry       string `json:"billing_country"`
	BillingAccountNumber string `json:"billing_account_number"`
}
