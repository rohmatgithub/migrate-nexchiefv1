package model

type NexsellerCustomer struct {
	ID                           int64
	CompanyProfileID             int64
	NcCode                       string `json:"principalID"`
	MnCode                       string `json:"distributorID"`
	Code                         string `json:"customerID"`
	Name                         string `json:"customerName"`
	Address1                     string `json:"customerAddress1"`
	Address2                     string `json:"customerAddress2"`
	Address3                     string `json:"customerAddress3"`
	City                         string `json:"customerCity"`
	Phone                        string `json:"customerPhone"`
	MsgNumber                    string `json:"customerMsgNumber"`
	Fax                          string `json:"customerFax"`
	Email                        string `json:"customerEmail"`
	AreaCode                     string `json:"areaID"`
	SubAreaCode                  string `json:"subareaID"`
	NexsellerAreaHierarchyID     int64
	MarketSegment                string `json:"marketSegmentID"`
	TypeCode                     string `json:"customerTypeID"`
	SubTypeCode                  string `json:"customerSubTypeID"`
	NexsellerChannelHierarchyID  int64
	NexsellerChanneHierarchyName string
	GroupCode                    string `json:"customerGroupID"`
	GroupID                      int64
	GroupName                    string
	CategoryCode                 string `json:"customerCategoryID"`
	CategoryName                 string
	CategoryID                   int64
	Class                        int     `json:"customerClass"`
	Status                       string  `json:"customerStatus"`
	IsBUMN                       string  `json:"customerIsBUMN"`
	IsPKP                        string  `json:"customerIsPKP"`
	Latitude                     float64 `json:"customerLatitude"`
	Longitude                    float64 `json:"customerLongitude"`
	ProvinceCode                 string  `json:"propinsiID"`
	ProvinceID                   int64
	ProvinceName                 string
	DistrictCode                 string `json:"kabupatenID"`
	DistrictID                   int64
	DistrictName                 string
	SubDistrictCode              string `json:"kecamatanID"`
	SubDistrictID                int64
	SubDistrictName              string
	UrbanVillageCode             string `json:"kelurahanID"`
	UrbanVillageID               int64
	UrbanVillageName             string
	FirstTransactionDate         string `json:"firstTransactionDate"`
	IsCustomerTrade              string `json:"isCustomerTrade"`
	IsPicos                      string `json:"isPicos"`
	PcosDate                     string `json:"pcosDate"`
	JoinDate                     string `json:"joinDate"`
	GromartFirstTransaction      string `json:"gromartFirstTransaction"`
	GromartLastTransaction       string `json:"gromartLastTransaction"`
	LastSync                     string `json:"lastSync"`
	IslandCode                   string `json:"islandID"`
	IslandID                     int64
	IslandName                   string
	StoreLocationCode            string `json:"storelocationID"`
	StoreLocationID              int64
	StoreLocationName            string
	StoreStatusCode              string `json:"storestatusID"`
	StoreStatusID                int64
	StoreStaturName              string
	LocationRemark               string `json:"locationRemark"`
	IsTdWeb                      string `json:"isTDWeb"`
	UserCategory1Code            string `json:"userCategory1ID"`
	UserCategory1ID              int64
	UserCategory1Name            string
	UserCategory2Code            string `json:"userCategory2ID"`
	UserCategory2ID              int64
	UserCategory2Name            string
	UserCategory3Code            string `json:"userCategory3ID"`
	UserCategory3ID              int64
	UserCategory3Name            string
	FlagVerified                 string `json:"flagVerified"`
	PkChecksum                   string
}
