package models

type ComplianceCertificates struct {
	Id                 int    `json:"id"`
	CarrierID          int    `json:"carrier_id"`
	CertificateDetails string `json:"certificate_details"`
}
