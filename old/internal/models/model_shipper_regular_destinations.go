package models

type Destinations struct {
	Id          int64  `json:"id"`
	ShipperID   int    `json:"shipper_id"`
	Destination string `json:"destination"`
}
