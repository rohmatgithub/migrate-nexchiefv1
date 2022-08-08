package dao

import (
	"database/sql"
	"rohmat.co.id/model"
	"time"
)

func GetIsland(db *sql.DB, ncID int64, code string) (result model.MasterModel, err model.ErrorModel) {
	query := "SELECT id, island_name FROM island WHERE island_id = $1 " +
		"AND nexchief_account_id = $2 "
	errs := db.QueryRow(query, code, ncID).Scan(&result.ID, &result.Name)
	if errs != nil && errs != sql.ErrNoRows {
		err = generateErrorModel(errs)
	}
	return
}

func InsertIsland(db *sql.DB, ncID int64, masterModel *model.IslandModel) (err model.ErrorModel) {
	query := "INSERT INTO island " +
		" (nexchief_account_id, island_id, island_name ) " +
		"VALUES " +
		"($1, $2, $3) " +
		"RETURNING id"
	param := []interface{}{ncID, masterModel.Code, masterModel.IslandName}
	_, errS := db.Exec(query, param...)
	if errS != nil && errS != sql.ErrNoRows {
		err = generateErrorModel(errS)
	}
	return
}

func InsertNexsellerIsland(tx *sql.Tx, ncID, mnID int64, schema string, masterModel *model.MasterModel) (resultID int64, err model.ErrorModel) {
	query := "INSERT INTO " + getSchema("nexseller_island", schema) +
		" (nexchief_account_id, mapping_nexseller_id, code, name ) " +
		"VALUES " +
		"($1, $2, $3, $4) " +
		"RETURNING id"
	param := []interface{}{ncID, mnID, masterModel.Code.String, masterModel.Name.String}
	errS := tx.QueryRow(query, param...).Scan(&resultID)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return
}

func UpdateNexsellerIsland(tx *sql.Tx, id int64, schema string, masterModel *model.MasterModel) (err model.ErrorModel) {
	query := "UPDATE " + getSchema("nexseller_island", schema) +
		" SET name = $1, updated_at = $2 WHERE id = $3"
	param := []interface{}{masterModel.Name.String, time.Now(), id}
	_, errS := tx.Exec(query, param...)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return
}
