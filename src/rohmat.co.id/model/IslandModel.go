package model

type IslandModel struct {
	ID               int64
	NcCode      string `json:"principalID"`
	Code         string `json:"islandID"`
	IslandName       string `json:"islandName"`
	IslandCreated    string `json:"islandCreated"`
	IslandModified   string `json:"islandModified"`
	IslandModifiedBy string `json:"islandModifiedBy"`
}
