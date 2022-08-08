package dao

import (
	"database/sql"
	"rohmat.co.id/model"
	"time"
)

func GetProvince(db *sql.DB, code string) (result model.MasterModel, err model.ErrorModel) {
	query := "SELECT id, code, name FROM province WHERE code = $1 "
	errs := db.QueryRow(query, code).Scan(&result.ID, &result.Code, &result.Name)
	if errs != nil && errs != sql.ErrNoRows {
		err = generateErrorModel(errs)
	}
	return
}

func GetProvinceByDistrict(db *sql.DB, districtCode string) (resultCode string, err model.ErrorModel) {
	var temp sql.NullString
	query := "SELECT p.code FROM district d " +
		"LEFT JOIN province d ON d.parent_id = p.id " +
		"WHERE d.code = $1 "
	errs := db.QueryRow(query, districtCode).Scan(&temp)
	if errs != nil && errs != sql.ErrNoRows {
		err = generateErrorModel(errs)
		return
	}
	resultCode = temp.String
	return
}

func GetNexsellerProvince(db *sql.DB, schema, provinceCode string, mnID int64) (resultID int64, resultName string, err model.ErrorModel) {
	if provinceCode == "" {
		return
	}
	var (
		tempID   sql.NullInt64
		tempName sql.NullString
	)
	query := "SELECT id, name FROM " + getSchema("nexseller_province", schema) +
		" WHERE mapping_nexseller_id = $1 AND code = $2"
	errs := db.QueryRow(query, mnID, provinceCode).Scan(&tempID, &tempName)
	if errs != nil && errs != sql.ErrNoRows {
		err = generateErrorModel(errs)
		return
	}
	resultID = tempID.Int64
	resultName = tempName.String
	return
}

func InsertNexsellerProvince(tx *sql.Tx, ncID, mnID int64, schema string, masterModel *model.MasterModel) (resultID int64, err model.ErrorModel) {
	query := "INSERT INTO " + getSchema("nexseller_province", schema) +
		" (mapping_nexseller_id, code, name ) " +
		"VALUES " +
		"($1, $2, $3) " +
		"RETURNING id"
	param := []interface{}{mnID, masterModel.Code.String, masterModel.Name.String}
	errS := tx.QueryRow(query, param...).Scan(&resultID)
	if errS != nil && errS != sql.ErrNoRows {
		err = generateErrorModel(errS)
	}
	return
}

func UpdateNexsellerProvince(tx *sql.Tx, id int64, schema string, masterModel *model.MasterModel) (err model.ErrorModel) {
	query := "UPDATE " + getSchema("nexseller_province", schema) +
		" SET name = $1, updated_at = $2 WHERE id = $3"
	param := []interface{}{masterModel.Name.String, time.Now(), id}
	_, errS := tx.Exec(query, param...)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return
}
