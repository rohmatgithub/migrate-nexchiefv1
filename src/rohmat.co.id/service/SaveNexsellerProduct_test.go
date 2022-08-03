package service

import (
	"database/sql"
	"encoding/json"
	"rohmat.co.id/dao"
	"rohmat.co.id/model"
	"testing"
	"time"
)


func TestInsertNexsellerProduct(t *testing.T) {
	path := "C:\\cdc-tools\\data sql\\product-distributor-ID01018901.json"
	StartReadFile(path, SaveNexsellerProduct, "nexseller product")
	time.Sleep(5 * time.Second)
}

func SaveNexsellerProduct(db *sql.DB, dataByte []byte)(errorModel model.ErrorModel) {
	var (
		nexchiefAccount model.NexchiefAccount
		data            model.NexsellerProduct
		mnID            int64
	)
	_ = json.Unmarshal(dataByte, &data)
	data.NcCode = "UCI"
	data.MnCode = "ID01018901"
	nexchiefAccount, errorModel = dao.GetNexchiefAccountID(db, data.NcCode)
	if errorModel.Error != nil {
		return
	}
	// get mapping nexseller
	mnID, errorModel = dao.GetMappingNexsellerID(db, nexchiefAccount.ID.Int64, data.MnCode)
	if errorModel.Error != nil {
		return
	}
	errorModel = dao.GetNexProductFK(db, nexchiefAccount.Schema.String, mnID, &data)
	if errorModel.Error != nil {
		return
	}
	if data.ID == 0 {
		// insert
		errorModel = dao.InsertNexsellerProduct(db, nexchiefAccount, mnID, &data)
		if errorModel.Error != nil {
			return
		}
	}
	return
}
