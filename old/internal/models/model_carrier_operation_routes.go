package models

type OperationalRoutes struct {
	Id               int    `json:"id"`
	CarrierID        int    `json:"carrier_id"`
	RouteDescription string `json:"route_description"`
	ServiceArea      string `json:"service_area"`
}
