package service

import (
	"database/sql"
	"encoding/json"
	"rohmat.co.id/config"
	"rohmat.co.id/dao"
	"rohmat.co.id/model"
)

func StartSaveIsland() {
	path := config.ApplicationConfiguration.GetDirPath().PathDir +
		config.ApplicationConfiguration.GetDirPath().Island

	StartReadFile(path, saveIsland, "island")
}

func saveIsland(db *sql.DB, dataByte []byte) (err model.ErrorModel) {
	var data model.IslandModel
	_ = json.Unmarshal(dataByte, &data)

	nexchiefAccount, err := dao.GetNexchiefAccountID(db, data.NcCode)
	if err.Error != nil {
		return
	}

	masterModel, err := dao.GetIsland(db, nexchiefAccount.ID.Int64, data.Code)
	if err.Error != nil {
		return
	} else if masterModel.ID.Int64 == 0 {
		// insert
		err = dao.InsertIsland(db, nexchiefAccount.ID.Int64, &data)
	}
	return
}
