package models

type BrokerAddress struct {
	ID       string `json:"id"`
	BrokerID int64  `json:"broker_id"`
	Street   string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
	ZipCode  string `json:"zip_code"`
	Country  string `json:"country"`
}
