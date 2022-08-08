package dao

import (
	"database/sql"
	"rohmat.co.id/model"
	"time"
)

func GetUserCategory(db *sql.DB, ncId int64, types int, schema string, data *model.NexsellerCustomer) (result model.MasterModel, err model.ErrorModel) {
	query := "SELECT id, code, name FROM " + getSchema("customer_category", schema) +
		" WHERE nexchief_account_id = $1 AND code = $2 AND type = $3"
	errs := db.QueryRow(query, ncId, data.CategoryCode, types).Scan(&result.ID, &result.Code, &result.Name)
	if errs != nil && errs != sql.ErrNoRows {
		err = generateErrorModel(errs)
	}
	return
}

func InsertNexsellerUserCategory(tx *sql.Tx, ncID, mnID int64, types int, schema string, masterModel *model.MasterModel) (resultID int64, err model.ErrorModel) {
	query := "INSERT INTO " + getSchema("nexseller_user_category", schema) +
		" (nexchief_account_id, mapping_nexseller_id, code, name, type ) " +
		"VALUES " +
		"($1, $2, $3, $4, $5) " +
		"RETURNING id"
	param := []interface{}{ncID, mnID, masterModel.Code.String, masterModel.Name.String, masterModel.ParentID.Int64, types}
	errS := tx.QueryRow(query, param...).Scan(&resultID)
	if errS != nil && errS != sql.ErrNoRows {
		err = generateErrorModel(errS)
	}
	return
}

func UpdateNexsellerUserCategory(tx *sql.Tx, id int64, schema string, masterModel *model.MasterModel) (err model.ErrorModel) {
	query := "UPDATE " + getSchema("nexseller_user_category", schema) +
		" SET name = $1, updated_at = $2 WHERE id = $3"
	param := []interface{}{masterModel.Name.String, time.Now(), id}
	_, errS := tx.Exec(query, param...)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return
}
