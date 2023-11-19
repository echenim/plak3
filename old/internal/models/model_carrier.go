package models

type Carrier struct {
	CarrierID          int    `json:"carrier_id"`
	CompanyName        string `json:"company_name"`
	ContactName        string `json:"contact_name"`
	ContactEmail       string `json:"contact_email"`
	ContactPhone       string `json:"contact_phone"`
	AddressID          int    `json:"address_id"`
	AvailabilityStatus string `json:"availability_status"`

	
}
