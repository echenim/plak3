package models

type TransportTypes struct {
	Id            int    `json:"id"`
	CarrierID     int    `json:"carrier_id"`
	TransportType string `json:"transport_type"`
}
