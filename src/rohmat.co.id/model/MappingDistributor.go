package model


type AddInfo struct {
	MappingField1  string `json:"mapping_field_1"`
	MappingField2  string `json:"mapping_field_2"`
	MappingField3  string `json:"mapping_field_3"`
	MappingField4  string `json:"mapping_field_4"`
	MappingField5  string `json:"mapping_field_5"`
	MappingField6  string `json:"mapping_field_6"`
	MappingField7  string `json:"mapping_field_7"`
	MappingField8  string `json:"mapping_field_8"`
	MappingField9  string `json:"mapping_field_9"`
	MappingField10 string `json:"mapping_field_10"`
}
type MappingNexseller struct {
	ID                  int64
	GeoTreeID           int64
	PriceGroupID        int64
	ProductCategoryID   int64
	NcCode              string `json:"principalID"`
	Code                string `json:"distributorID"`
	IsProductMapping    string `json:"isProductMapping"`
	IsSalesmanMapping   string `json:"isSalesmanMapping"`
	IsCustomerMapping   string `json:"isCustomerMapping"`
	UserLevel1          string `json:"userLevel1ID"`
	UserLevel2          string `json:"userLevel2ID"`
	UserLevel3          string `json:"userLevel3ID"`
	UserLevel4          string `json:"userLevel4ID"`
	UserLevel5          string `json:"userLevel5ID"`
	MappingField1       string `json:"mappingField1"`
	MappingField2       string `json:"mappingField2"`
	MappingField3       string `json:"mappingField3"`
	MappingField4       string `json:"mappingField4"`
	MappingField5       string `json:"mappingField5"`
	MappingField6       string `json:"mappingField6"`
	MappingField7       string `json:"mappingField7"`
	MappingField8       string `json:"mappingField8"`
	MappingField9       string `json:"mappingField9"`
	MappingField10      string `json:"mappingField10"`
	ResignDate          string `json:"resignDate"`
	ActiveDate          string `json:"activeDate"`
	PriceGroupCode      string `json:"priceGroupID"`
	ProductClassFrom    int    `json:"productClassFrom"`
	ProductClassThru    int    `json:"productClassThru"`
	ProductCategoryCode string `json:"productCategoryID"`
	EmailData           string `json:"emailData"`
	SyncMethod          string `json:"syncMethod"`
	LastDMSSyncStr      string `json:"lastDMSSync"`
	LastSFASyncStr      string `json:"lastSFASync"`
	//LastDMSSync             time.Time
	//LastSFASync             time.Time
	HostingOnly             string `json:"hostingOnly"`
	GromartMerchantID       string `json:"gromartMerchantID"`
	IsGenerateDelivery      string `json:"isGenerateDelivery"`
	EmailTo                 string `json:"emailTo"`
	EmailToCC               string `json:"emailToCC"`
	PrefixDeleted           string `json:"prefixDeleted"`
	Nd6ClosedDate           string `json:"nd6ClosedDate"`
	SocketUserID            string `json:"socketUserID"`
	SocketPassword          string `json:"socketPassword"`
	SocketStatus            string `json:"socketStatus"`
	ProductRegistrationDate string `json:"productRegistrationDate"`
	ProductValidThru        string `json:"productValidThru"`
}

type Distributor struct {
	ID         int64
	Code       string `json:"distributorID"`
	Name       string `json:"distributorName"`
	Address1   string `json:"address1"`
	Address2   string `json:"address2"`
	City       string `json:"city"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Fax        string `json:"fax"`
	ParentCode string `json:"parentID"`
	Npwp       string `json:"distributorNPWP"`
}

type MappingDistributor struct {
	NcID int64
	ID   int64
}
