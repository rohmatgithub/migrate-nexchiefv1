package model

type VendorModel struct {
	ID                 int64
	PrincipalID        string `json:"principalID"`
	DistributorID      string `json:"distributorID"`
	Code               string `json:"vendorID"`
	Name               string `json:"vendorName"`
	NexchiefAccountID  int64
	MappingPrincipalID int64
	CompanyProfileID   int64
}
