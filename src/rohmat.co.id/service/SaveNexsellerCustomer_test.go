package service

import (
	"database/sql"
	"encoding/json"
	"rohmat.co.id/dao"
	"rohmat.co.id/model"
	"testing"
	"time"
)

//func TestRead(t *testing.T){
//	var (
//		temp            structNexsellerCustomer
//	)
//	errorModel := gonfig.GetConf("C:\\cdc-tools\\data sql\\customer-join-promotion-budget-customer.json", &temp)
//	if errorModel != nil {
//		assert.FailNow(t, errorModel.Error())
//	}
//}
func TestSaveNexsellerCustomer(t *testing.T) {
	// 2021-01-01
	path := "C:\\cdc-tools\\data sql\\sbp\\customer_distributor_218039.json"
	StartReadFile(path, SaveNexsellerCustomer)
	time.Sleep(5 * time.Second)
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

//func TestStart(t *testing.T) {
//	path := "C:\\cdc-tools\\data sql\\test.json"
//	StartReadFile(path, run)
//}
//
//func run(db *sql.DB, data []byte) model.ErrorModel {
//	var customer model.NexsellerCustomer
//	_ = json.Unmarshal(data, &customer)
//	log.Println(">>", customer.Code)
//	return model.ErrorModel{}
//}
