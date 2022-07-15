package service

import (
	"database/sql"
	"encoding/json"
	"rohmat.co.id/dao"
	"rohmat.co.id/model"
	"testing"
)

func TestInsertSalesman(t *testing.T) {
	path := "C:\\cdc-tools\\data sql\\local\\salesman.json"
	StartReadFile(path, SaveSalesman)
	//time.Sleep(5 * time.Second)
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
		err = dao.InsertSalesman(db, &data, nexchiefAccount, personProfileID, mnID)
		//err = insertNexsellerProduct(db, nexchiefAccount, mnID, &data)
		if err.Error != nil {
			return
		}
	}
	return
}
