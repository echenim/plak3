package models

type FleetDetails struct {
	ID                  int    `json:"id"`
	CarrierID           int    `json:"carrier_id"`
	Vehicle             string `json:"vehicle"`
	Capacity            int    `json:"capacity"`
	MaintenanceSchedule string `json:"maintenance_schedule"`
}
