package service

import (
	"database/sql"
	"encoding/json"
	"rohmat.co.id/config"
	"rohmat.co.id/dao"
	"rohmat.co.id/model"
)

func StartSaveNexsellerCustomer() {
	// 2021-01-01
	path := config.ApplicationConfiguration.GetDirPath().PathDir +
		config.ApplicationConfiguration.GetDirPath().Customer
	StartReadFile(path, SaveNexsellerCustomer, "nexseller customer")
}

func SaveNexsellerCustomer(db *sql.DB, dataByte []byte) (errorModel model.ErrorModel) {
	var (
		nexchiefAccount model.NexchiefAccount
		data            model.NexsellerCustomer
		mnID int64
	)
	_ = json.Unmarshal(dataByte, &data)

	nexchiefAccount, errorModel = dao.GetNexchiefAccountID(db, data.NcCode)
	if errorModel.Error != nil {
		return
	}
	// get mapping nexseller
	mnID, errorModel = dao.GetMappingNexsellerID(db, nexchiefAccount.ID.Int64, data.MnCode)
	if errorModel.Error != nil {
		return
	}
	// get data FK
	errorModel = dao.GetNexsellerCustomerFK(db, nexchiefAccount, &data, mnID)
	if errorModel.Error != nil {
		return
	} else if data.ID == 0 {
		// get company profile id
		errorModel = dao.GetCompanyProfile(db, &data)
		if errorModel.Error != nil {
			return
		}
		errorModel = dao.InsertNexsellerCustomer(db, nexchiefAccount, mnID, &data)
		if errorModel.Error != nil {
			return
		}
	}
	return
}
