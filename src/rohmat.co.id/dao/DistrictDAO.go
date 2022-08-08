package dao

import (
	"database/sql"
	"rohmat.co.id/model"
	"time"
)

func GetDistrict(db *sql.DB, code string) (result model.MasterModel, err model.ErrorModel) {
	query := "SELECT id, code, name FROM district WHERE code = $1 "
	errs := db.QueryRow(query, code).Scan(&result.ID, &result.Code, &result.Name)
	if errs != nil && errs != sql.ErrNoRows {
		err = generateErrorModel(errs)
	}
	return
}

func GetDistrictBySubDistrict(db *sql.DB, subDistrictCode string) (resultCode string, err model.ErrorModel) {
	var temp sql.NullString
	query := "SELECT d.code FROM sub_district sd " +
		"LEFT JOIN district d ON sd.parent_id = d.id " +
		"WHERE sd.code = $1 "
	errs := db.QueryRow(query, subDistrictCode).Scan(&temp)
	if errs != nil && errs != sql.ErrNoRows {
		err = generateErrorModel(errs)
		return
	}
	resultCode = temp.String
	return
}

func GetNexsellerDistrict(db *sql.DB, schema, districtCode string, mnID int64) (resultID int64, resultName string, err model.ErrorModel) {
	if districtCode == "" {
		return
	}
	var (
		tempID   sql.NullInt64
		tempName sql.NullString
	)
	query := "SELECT id, name FROM " + getSchema("nexseller_district", schema) +
		" WHERE mapping_nexseller_id = $1 AND code = $2"
	errs := db.QueryRow(query, mnID, districtCode).Scan(&tempID, &tempName)
	if errs != nil && errs != sql.ErrNoRows {
		err = generateErrorModel(errs)
		return
	}
	resultID = tempID.Int64
	resultName = tempName.String
	return
}

func InsertNexsellerDistrict(tx *sql.Tx, ncID, mnID int64, schema string, masterModel *model.MasterModel) (resultID int64, err model.ErrorModel) {
	query := "INSERT INTO " + getSchema("nexseller_district", schema) +
		" (mapping_nexseller_id, code, name, parent_id ) " +
		"VALUES " +
		"($1, $2, $3, $4) " +
		"RETURNING id"
	param := []interface{}{mnID, masterModel.Code.String, masterModel.Name.String, masterModel.ParentID.Int64}
	errS := tx.QueryRow(query, param...).Scan(&resultID)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return
}

func UpdateNexsellerDistrict(tx *sql.Tx, id int64, schema string, masterModel *model.MasterModel) (err model.ErrorModel) {
	query := "UPDATE " + getSchema("nexseller_district", schema) +
		" SET name = $1, updated_at = $2 WHERE id = $3"
	param := []interface{}{masterModel.Name.String, time.Now(), id}
	_, errS := tx.Exec(query, param...)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return
}
