package dao

import (
	"database/sql"
	"rohmat.co.id/model"
	"time"
)

func GetUrbanVillage(db *sql.DB, code string) (result model.MasterModel, err model.ErrorModel) {
	query := "SELECT id, code, name FROM sub_district WHERE code = $1 "
	errs := db.QueryRow(query, code).Scan(&result.ID, &result.Code, &result.Name)
	if errs != nil && errs != sql.ErrNoRows {
		err = generateErrorModel(errs)
	}
	return
}

func InsertNexsellerUrbanVillage(tx *sql.Tx, ncID, mnID int64, schema string, masterModel *model.MasterModel) (resultID int64, err model.ErrorModel) {
	query := "INSERT INTO " + getSchema("nexseller_urban_village", schema) +
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

func UpdateNexsellerUrbanVillage(tx *sql.Tx, id int64, schema string, masterModel *model.MasterModel) (err model.ErrorModel) {
	query := "UPDATE " + getSchema("nexseller_urban_village", schema) +
		" SET name = $1, updated_at = $2 WHERE id = $3"
	param := []interface{}{masterModel.Name.String, time.Now(), id}
	_, errS := tx.Exec(query, param...)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return
}
