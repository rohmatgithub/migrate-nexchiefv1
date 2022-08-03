package dao

import (
	"database/sql"
	"rohmat.co.id/model"
)

func GetMappingPrincipalID(db *sql.DB, ncID int64, code string) (result int64, err model.ErrorModel) {
	query := "SELECT id FROM mapping_principal where code = $1 AND nexchief_account_id = $2 "
	errS := db.QueryRow(query, code, ncID).Scan(&result)
	if errS != nil && errS != sql.ErrNoRows {
		err = generateErrorModel(errS)
	}
	return
}

func GetDivisionFK(db *sql.DB, model *model.DivisionModel) (err model.ErrorModel) {
	query := "SELECT id FROM division WHERE nexchief_account_id = $1 AND mapping_principal_id = $2 " +
		"AND code = $3 "
	errS := db.QueryRow(query, model.NexchiefAccountID, model.MappingPrincipalID, model.Code).Scan(&model.ID)
	if errS != nil && errS != sql.ErrNoRows {
		err = generateErrorModel(errS)
	}
	return
}

func InsertDivision(db *sql.DB, model *model.DivisionModel) (err model.ErrorModel) {
	query := "INSERT INTO division " +
		"(nexchief_account_id, name, code, mapping_principal_id ) " +
		"VALUES " +
		"($1, $2, $3, $4 ) "
	param := []interface{}{model.NexchiefAccountID, model.Name, model.Code, model.MappingPrincipalID}
	_, errS := db.Exec(query, param...)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return
}
