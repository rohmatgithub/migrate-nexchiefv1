package service

import (
	"database/sql"
	"encoding/json"
	"rohmat.co.id/config"
	"rohmat.co.id/dao"
	"rohmat.co.id/model"
)

func StartSaveDivision() {
	path := config.ApplicationConfiguration.GetDirPath().PathDir +
		config.ApplicationConfiguration.GetDirPath().Division

	StartReadFile(path, SaveDivision, "division")
}
func SaveDivision(db *sql.DB, dataByte []byte)  (err model.ErrorModel) {
	var (
		data            model.DivisionModel
	)
	_ = json.Unmarshal(dataByte, &data)

	nexchiefAccount, err := dao.GetNexchiefAccountID(db, data.PrincipalID)
	if err.Error != nil {
		return
	}
	data.NexchiefAccountID = nexchiefAccount.ID.Int64

	data.MappingPrincipalID, err = dao.GetMappingPrincipalID(db, nexchiefAccount.ID.Int64, data.PrincipalID)
	if err.Error != nil {
		return
	}

	err = dao.GetDivisionFK(db, &data)
	if err.Error != nil {
		return
	} else if data.ID == 0 {
		// insert

		err = dao.InsertDivision(db, &data)
		if err.Error != nil {
			return
		}
	}
	return
}

