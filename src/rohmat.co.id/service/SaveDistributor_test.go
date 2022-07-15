package service

import (
	"database/sql"
	"encoding/json"
	"rohmat.co.id/dao"
	"rohmat.co.id/model"
	"testing"
	"time"
)


func TestInsertCompanyProfile(t *testing.T) {
	path := "C:\\cdc-tools\\data sql\\local\\distributor.json"

	StartReadFile(path, SaveDistributor)
	time.Sleep(5 * time.Second)
}

func SaveDistributor(db *sql.DB, dataByte []byte)  (errorModel model.ErrorModel) {
	var (
		cpID, parentID int64
		resultDB       []model.MappingDistributor
		data            model.Distributor
	)
	_ = json.Unmarshal(dataByte, &data)

	cpID, errorModel = dao.GetIDCompanyProfile(db, &data)
	if errorModel.Error != nil {
		return
	} else if cpID == 0 {
		// insert
		cpID, errorModel = dao.InsertCompanyProfile(db, &data)
		if errorModel.Error != nil {
			return
		}
	}
	resultDB, errorModel = dao.GetDistributor(db, &data)
	if errorModel.Error != nil {
		return
	}

	for _, m := range resultDB {
		parentID, errorModel = dao.GetParentID(db, &m)
		if errorModel.Error != nil {
			return
		}
		errorModel = dao.UpdateMappingNexseller(db, m.ID, cpID, parentID)
		if errorModel.Error != nil {
			return
		}
	}
	return
}
