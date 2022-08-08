package dao

import (
	"database/sql"
	"rohmat.co.id/model"
	"time"
)

func GetCustomerGroup(db *sql.DB, ncId int64, data *model.NexsellerCustomer) (result model.MasterModel, err model.ErrorModel) {
	query := "SELECT id, customer_group_id, customer_group_name FROM customer_group " +
		"WHERE nexchief_account_id = $1 AND customer_group_id = $2 "
	errs := db.QueryRow(query, ncId, data.GroupCode).Scan(&result.ID, &result.Code, &result.Name)
	if errs != nil && errs != sql.ErrNoRows {
		err = generateErrorModel(errs)
	}
	return
}

func InsertNexsellerCustomerGroup(tx *sql.Tx, ncID, mnID int64, schema string, masterModel *model.MasterModel) (resultID int64, err model.ErrorModel) {
	query := "INSERT INTO " + getSchema("nexseller_customer_group", schema) +
		" (nexchief_account_id, mapping_nexseller_id, customer_group_id, customer_group_name ) " +
		"VALUES " +
		"($1, $2, $3, $4) " +
		"RETURNING id"
	param := []interface{}{ncID, mnID, masterModel.Code.String, masterModel.Name.String}
	errS := tx.QueryRow(query, param...).Scan(&resultID)
	if errS != nil && errS != sql.ErrNoRows {
		err = generateErrorModel(errS)
	}
	return
}

func UpdateNexsellerCustomerGroup(tx *sql.Tx, id int64, schema string, masterModel *model.MasterModel) (err model.ErrorModel) {
	query := "UPDATE " + getSchema("nexseller_customer_group", schema) +
		" SET customer_group_name = $1, updated_at = $2 WHERE id = $3"
	param := []interface{}{masterModel.Name.String, time.Now(), id}
	_, errS := tx.Exec(query, param...)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return
}
