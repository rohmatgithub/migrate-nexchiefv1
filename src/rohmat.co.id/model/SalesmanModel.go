package model

type Salesman struct {
	ID                      int64
	NcCode                  string `json:"principalID"`
	MnCode                  string `json:"distributorID"`
	Code                    string `json:"salesmanID"`
	Name                    string `json:"salesmanName"`
	Address1                string `json:"salesmanAddress1"`
	Address2                string `json:"salesmanAddress2"`
	City                    string `json:"salesmanCity"`
	Phone                   string `json:"salesmanPhone"`
	Email                   string `json:"salesmanEmail"`
	Nik                     string `json:"salesmanKTP"`
	JoinDate                string `json:"salesmanJoinDate"`
	ResignDate              string `json:"salesmanResignDate"`
	Status                  string `json:"isSalesmanActive"`
	Type                    string `json:"salesmanType"`
	NexileOn                string `json:"salesmanNexileOn"`
	LastSync                string `json:"lastSync"`
	ImeiNumber              string `json:"imeiNumber"`
	ImeiNumber2             string `json:"imeiNumber2"`
	RegistrationDate        string `json:"registrationDate"`
	RegistrationCity        string `json:"registrationCity"`
	RegisLat                string `json:"registrationLat"`
	RegisLong               string `json:"registrationLong"`
	Category                string `json:"salesmanCategory"`
	PrincipalSalesmanType   string `json:"principalSalesmanType"`
	PrincipalSalesmanTypeID int64
	Group                   string `json:"salesmanGroup"`
	NexmileVersion          string `json:"nexmileVersion"`
	NexmileValidThru        string `json:"nexmileValidThru"`
	NexmileDeviceID         string `json:"nexmileDeviceID"`
}

