package model

type DivisionModel struct {
	ID                 int64
	PrincipalID        string `json:"principalID"`
	DistributorID      string `json:"distributorID"`
	Code               string `json:"divisionID"`
	Name               string `json:"divisionName"`
	NexchiefAccountID  int64
	MappingPrincipalID int64
}
