package dao

import (
	"database/sql"
	"rohmat.co.id/model"
	"time"
)

func GetSubDistrict(db *sql.DB, code string) (result model.MasterModel, err model.ErrorModel) {
	query := "SELECT id, code, name FROM sub_district WHERE code = $1 "
	errs := db.QueryRow(query, code).Scan(&result.ID, &result.Code, &result.Name)
	if errs != nil && errs != sql.ErrNoRows {
		err = generateErrorModel(errs)
	}
	return
}

func GetSubDistrictByUrbanVillage(db *sql.DB, urbanVillageCode string) (resultCode string, err model.ErrorModel) {
	var temp sql.NullString
	query := "SELECT sd.code FROM urban_village uv " +
		"LEFT JOIN sub_district sd ON uv.parent_id = sd.id " +
		"WHERE uv.code = $1 "
	errs := db.QueryRow(query, urbanVillageCode).Scan(&temp)
	if errs != nil && errs != sql.ErrNoRows {
		err = generateErrorModel(errs)
		return
	}
	resultCode = temp.String
	return
}

func GetNexsellerSubDistrict(db *sql.DB, schema, subDistrictCode string, mnID int64) (resultID int64, resultName string, err model.ErrorModel) {
	if subDistrictCode == "" {
		return
	}
	var (
		tempID   sql.NullInt64
		tempName sql.NullString
	)
	query := "SELECT id, name FROM " + getSchema("nexseller_sub_district", schema) +
		" WHERE mapping_nexseller_id = $1 AND code = $2"
	errs := db.QueryRow(query, mnID, subDistrictCode).Scan(&tempID, &tempName)
	if errs != nil && errs != sql.ErrNoRows {
		err = generateErrorModel(errs)
		return
	}
	resultID = tempID.Int64
	resultName = tempName.String
	return
}

func InsertNexsellerSubDistrict(tx *sql.Tx, ncID, mnID int64, schema string, masterModel *model.MasterModel) (resultID int64, err model.ErrorModel) {
	query := "INSERT INTO " + getSchema("nexseller_sub_district", schema) +
		" (mapping_nexseller_id, code, name, parent_id ) " +
		"VALUES " +
		"($1, $2, $3, $4) " +
		"RETURNING id"
	param := []interface{}{mnID, masterModel.Code.String, masterModel.Name.String, masterModel.ParentID.Int64}
	errS := tx.QueryRow(query, param...).Scan(&resultID)
	if errS != nil && errS != sql.ErrNoRows {
		err = generateErrorModel(errS)
	}
	return
}

func UpdateNexsellerSubDistrict(tx *sql.Tx, id int64, schema string, masterModel *model.MasterModel) (err model.ErrorModel) {
	query := "UPDATE " + getSchema("nexseller_sub_district", schema) +
		" SET name = $1, updated_at = $2 WHERE id = $3"
	param := []interface{}{masterModel.Name.String, time.Now(), id}
	_, errS := tx.Exec(query, param...)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return
}
