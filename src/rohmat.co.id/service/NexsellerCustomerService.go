package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nexsoft.co.id/nexcommon/util"
	"rohmat.co.id/config"
	"rohmat.co.id/dao"
	"rohmat.co.id/model"
)

func StartSaveNexsellerCustomer() {
	// 2021-01-01
	path := config.ApplicationConfiguration.GetDirPath().PathDir +
		config.ApplicationConfiguration.GetDirPath().Customer

	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i, file := range files {
		if !file.IsDir() {
			StartReadFile(path+"/"+file.Name(), SaveNexsellerCustomer, fmt.Sprintf("nexseller customer %d", i))
		}
	}
}

func SaveNexsellerCustomer(db *sql.DB, dataByte []byte) (errorModel model.ErrorModel) {
	var (
		nexchiefAccount model.NexchiefAccount
		data            model.NexsellerCustomer
		mnID            int64
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
		pkString := data.NcCode + data.MnCode + data.Code
		data.PkChecksum = util.CheckSumWithXXHASH([]byte(pkString))
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
