package service

import (
	"database/sql"
	"encoding/json"
	"rohmat.co.id/dao"
	"rohmat.co.id/model"
	"testing"
	"time"
)

func TestInsertMappingNexseller(t *testing.T) {
	path := "C:\\cdc-tools\\data sql\\local\\mapping-distributor.json"

	StartReadFile(path, saveMappingDistributor)
	time.Sleep(5 * time.Second)
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