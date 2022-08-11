package dao

import (
	"database/sql"
	"rohmat.co.id/model"
)

func GetListIDNexchiefAccountForScope(db *sql.DB) (resultList []int64, err model.ErrorModel) {
	query := "SELECT id FROM nexchief_account order by id"
	rows, errS := db.Query(query)
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
			var id sql.NullInt64

			errS = rows.Scan(&id)
			if errS != nil {
				err = generateErrorModel(errS)
				return
			} else if id.Int64 > 0 {
				resultList = append(resultList, id.Int64)
			}
		}
	}
	return
}

func InsertDataScope(db *sql.DB, model model.ScopeModel) (err model.ErrorModel) {
	query := "INSERT INTO data_scope " +
		"(scope, description, created_by, updated_by) " +
		"VALUES " +
		"($1, $2, $3, $4) " +
		"ON CONFLICT(scope) DO NOTHING;"
	_, errs := db.Exec(query, model.Scope, model.Description, model.CreatedBy, model.CreatedBy)
	if errs != nil {
		err = generateErrorModel(errs)
	}
	return
}
