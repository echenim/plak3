package models

type Shipper struct {
	ShipperID    int            `json:"shipper_id"`
	CompanyName  string         `json:"company_name"`
	ContactName  string         `json:"contact_name"`
	ContactEmail string         `json:"contact_email"`
	ContactPhone string         `json:"contact_phone"`
	Address      ShipperAddress `json:"address"`
	// InventoryManagementSystemID *int
	// CreditRating                string

	RegularDestinations []Destinations
	ContractedBrokers   []ContractedShipperBrokers
	BillingInformation  BillingInformation
}
