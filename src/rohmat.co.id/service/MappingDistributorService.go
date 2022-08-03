package service

import (
	"database/sql"
	"encoding/json"
	"rohmat.co.id/config"
	"rohmat.co.id/dao"
	"rohmat.co.id/model"
)

//dijalankan sebelum save distributor  StartSaveDistributor()
func StartSaveMappingDistributor() {
	path := config.ApplicationConfiguration.GetDirPath().PathDir +
		config.ApplicationConfiguration.GetDirPath().MappingDistributor

	StartReadFile(path, saveMappingDistributor, "mapping distributor")
}

func saveMappingDistributor(db *sql.DB, dataByte []byte) (err model.ErrorModel) {
	var data model.MappingNexseller
	_ = json.Unmarshal(dataByte, &data)

	nexchiefAccount, err := dao.GetNexchiefAccountID(db, data.NcCode)
	if err.Error != nil {
		return
	}

	err = dao.GetMappingNexseller(db, nexchiefAccount.ID.Int64, &data)
	if err.Error != nil {
		return
	} else if data.ID == 0 {
		// insert
		err = dao.InsertMappingNexseller(db, nexchiefAccount.ID.Int64, &data)
	}
	if err.Error != nil {
		return
	}
	return
}