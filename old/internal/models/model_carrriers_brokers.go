package models

type ContractedCarrierBrokers struct {
	Id        int64 `json:"id"`
	CarrierID int   `json:"carrier_id"`
	BrokerID  int   `json:"broker_id"`
}
