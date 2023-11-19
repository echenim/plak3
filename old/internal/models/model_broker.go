package models

type Broker struct {
	BrokerID            int           `json:"broker_id"`
	CompanyName         string        `json:"company_name"`
	ContactName         string        `json:"contact_name"`
	ContactEmail        string        `json:"contact_email"`
	ContactPhone        string        `json:"contact_phone"`
	Address             BrokerAddress `json:"address"`
	ServicePortfolio    string        `json:"service_portfolio"`
	ClientList          string        `json:"client_list"`
	CommissionStructure string        `json:"commission_structure"`
	LicenseInformation  string        `json:"license_information"`
	PerformanceMetrics  string        `json:"performance_metrics"`

	ContractedBrokers []ContractedShipperBrokers
}
