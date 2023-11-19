package models

type ContractedShipperBrokers struct {
	Id        int64 `json:"id"`
	ShipperID int64 `json:"shipper_id"`
	BrokerID  int64 `json:"broker_id"`
}
