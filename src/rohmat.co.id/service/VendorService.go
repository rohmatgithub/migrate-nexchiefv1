package service

import (
	"database/sql"
	"encoding/json"
	"rohmat.co.id/dao"
	"rohmat.co.id/model"
)
func StartSaveVendor() {
	path := "C:\\cdc-tools\\data sql\\local\\vendor.json"

	StartReadFile(path, SaveVendor, "vendor")
}

func SaveVendor(db *sql.DB, dataByte []byte)  (err model.ErrorModel) {
	var (
		data            model.VendorModel
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

	err = dao.GetVendorFK(db, &data, nexchiefAccount.Schema.String)
	if err.Error != nil {
		return
	}

	if data.CompanyProfileID == 0 {
		err = dao.InsertCompanyFromVendor(db, &data)
		if err.Error != nil {
			return
		}
	}
	if data.ID == 0 {
		// insert
		err = dao.InsertVendor(db, &data, nexchiefAccount.Schema.String)
		if err.Error != nil {
			return
		}
	}
	return
}

