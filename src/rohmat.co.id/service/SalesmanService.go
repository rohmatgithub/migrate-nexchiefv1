package service

import (
	"database/sql"
	"encoding/json"
	"nexsoft.co.id/nexcommon/util"
	"rohmat.co.id/config"
	"rohmat.co.id/dao"
	"rohmat.co.id/model"
)

func StartInsertSalesman() {
	path := config.ApplicationConfiguration.GetDirPath().PathDir +
		config.ApplicationConfiguration.GetDirPath().Salesman
	StartReadFile(path, SaveSalesman, "salesman")
}

func SaveSalesman(db *sql.DB, dataByte []byte) (err model.ErrorModel) {
	var (
		nexchiefAccount       model.NexchiefAccount
		data            model.Salesman
		mnID, personProfileID int64
	)
	_ = json.Unmarshal(dataByte, &data)

	nexchiefAccount, err = dao.GetNexchiefAccountID(db, data.NcCode)
	if err.Error != nil {
		return
	}
	// get mapping nexseller
	mnID, err = dao.GetMappingNexsellerID(db, nexchiefAccount.ID.Int64, data.MnCode)
	if err.Error != nil {
		return
	}
	// get data FK
	err = dao.GetSalesmanFKID(db, mnID, &data, nexchiefAccount.Schema.String)
	if err.Error != nil {
		return
	} else if data.ID == 0 {
		// get person profile id
		//personProfileID, err = dao.GetPersonProfileID(db, &data)
		//if err.Error != nil {
		//	return
		//}
		//else if personProfileID == 0 {
		//	personProfileID, err = insertPersonProfile(db, &data)
		//	if err.Error != nil {
		//		fmt.Println("err insert profile :", err.Error())
		//		return
		//	}
		//}
		pkString := data.NcCode + data.MnCode + data.Code
		data.PkChecksum = util.CheckSumWithXXHASH([]byte(pkString))
		err = dao.InsertSalesman(db, &data, nexchiefAccount, personProfileID, mnID)
		//err = insertNexsellerProduct(db, nexchiefAccount, mnID, &data)
		if err.Error != nil {
			return
		}
	}
	return
}
