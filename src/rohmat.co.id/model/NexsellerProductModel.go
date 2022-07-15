package model


type NexsellerProduct struct {
	ID                   int64
	NcCode               string `json:"principalID"`
	MnCode               string `json:"distributorID"`
	Code                 string `json:"productCode"`
	PrincipalProductCode string `json:"principalProductCode"`
	PrincipalProductID   int64
	Name                 string  `json:"productName"`
	Packaging            string  `json:"productPackaging"`
	Uom1                 string  `json:"uom1"`
	Uom2                 string  `json:"uom2"`
	Uom3                 string  `json:"uom3"`
	Uom4                 string  `json:"uom4"`
	Conversion1to4       int     `json:"conversion1to4"`
	Conversion2to4       int     `json:"conversion2to4"`
	Conversion3to4       int     `json:"conversion3to4"`
	Status               string  `json:"productStatus"`
	BuyingPrice          float64 `json:"buyingPrice"`
	SellingPrice         float64 `json:"sellingPrice"`
	NexchiefRatio        int     `json:"nexchiefRatio"`
	DivisionCode         string  `json:"divisionID"`
	DivisionID           int64
	VendorCode           string `json:"vendorID"`
	VendorID             int64
}
