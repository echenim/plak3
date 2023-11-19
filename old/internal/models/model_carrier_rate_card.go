package models

type RateCard struct {
	Id                 int64   `json:"id"`
	RateCardID         int     `json:"rate_card_id"`
	CarrierID          int     `json:"carrier_id"`
	ServiceDescription string  `json:"service_description"`
	Rate               float64 `json:"rate"` // Assuming DECIMAL(10, 2) maps to float64
}
