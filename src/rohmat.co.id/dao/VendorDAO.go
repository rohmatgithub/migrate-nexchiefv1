package dao

import (
	"database/sql"
	"rohmat.co.id/model"
)

func GetVendorFK(db *sql.DB,  data *model.VendorModel, schema string) (err model.ErrorModel) {
	query := "SELECT 'vendor', id from " + getSchema("vendor", schema) + " WHERE mapping_principal_id = $1 AND code = $2 " +
		"UNION ALL " +
		"SELECT 'cp', id FROM company_profile WHERE name = $3 "
	param := []interface{}{data.MappingPrincipalID, data.Code,
		data.Name}
	rows, errS := db.Query(query, param...)
	if errS != nil {
		err = generateErrorModel(errS)
		return
	}
	if rows != nil {
		defer func() {
			errS = rows.Close()
			if errS != nil {
				err = generateErrorModel(errS)
				return
			}
		}()

		for rows.Next() {
			var key string
			var id sql.NullInt64

			errS = rows.Scan(&key, &id)
			if errS != nil {
				err = generateErrorModel(errS)
				return
			}
			switch key {
			case "vendor":
				data.ID = id.Int64
			case "cp":
				data.CompanyProfileID = id.Int64
			}
		}
	}
	return
}

func InsertVendor(db *sql.DB, model *model.VendorModel, schema string) (err model.ErrorModel) {
	query := "INSERT INTO " + getSchema("vendor", schema) +
		"(nexchief_account_id, code, mapping_principal_id, " +
		"company_profile_id ) " +
		"VALUES " +
		"($1, $2, $3, $4 ) "
	param := []interface{}{model.NexchiefAccountID, model.Code, model.MappingPrincipalID, model.CompanyProfileID}
	_, errS := db.Exec(query, param...)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return
}

func InsertCompanyFromVendor(db *sql.DB, model *model.VendorModel) (err model.ErrorModel) {
	query := "INSERT INTO company_profile " +
		"(name ) " +
		"VALUES " +
		"($1) RETURNING id"
	param := []interface{}{model.Name}
	errS := db.QueryRow(query, param...).Scan(&model.CompanyProfileID)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return
}
