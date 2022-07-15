package model

type DataUserLevel struct {
	ID       int64
	ParentID int64
	Level    int
	Code     string
	Code1    string `json:"userLevel1ID"`
	Code2    string `json:"userLevel2ID"`
	Code3    string `json:"userLevel3ID"`
	Code4    string `json:"userLevel4ID"`
	Code5    string `json:"userLevel5ID"`
	NcCode   string `json:"principalID"`
	Name     string `json:"userLevelName"`
}
