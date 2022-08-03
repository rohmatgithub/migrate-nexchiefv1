package dao

import (
	"database/sql"
	"rohmat.co.id/model"
)

func InsertGeoTree(db *sql.DB, ncID int64, userParam *model.DataUserLevel) (err model.ErrorModel) {
	//funcName := "InsertDataGroup"
	query :=
		" INSERT INTO  geo_tree " +
			"  (nexchief_account_id, code, name," +
			"  level, parent_id, geo_tree_node ) " +
			" VALUES " +
			"($1, $2, $3, " +
			" $4, $5, $6) "
	param := []interface{}{
		ncID, userParam.Code, userParam.Name,
		userParam.Level, userParam.ParentID, userParam.Node}

	_, errS := db.Exec(query, param...)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return
}

func UpdateGeoTree(db *sql.DB, ncID int64, userParam *model.DataUserLevel) (err model.ErrorModel) {
	//funcName := "InsertDataGroup"
	query :=
		" UPDATE  geo_tree SET " +
			"nexchief_account_id = $1, code = $2, name = $3," +
			"level = $4, parent_id = $5, geo_tree_node = $6 " +
			"WHERE id = $7 "
	param := []interface{}{
		ncID, userParam.Code, userParam.Name,
		userParam.Level, userParam.ParentID, userParam.Node,
		userParam.ID}

	_, errS := db.Exec(query, param...)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return
}

func GetGeoTreeID(db *sql.DB, ncID int64, data *model.DataUserLevel) (err model.ErrorModel) {
	query := "SELECT id FROM geo_tree WHERE nexchief_account_id = $1 AND code = $2 " +
		"AND level = $3 AND parent_id = $4 "
	param := []interface{}{ncID, data.Code, data.Level, data.ParentID}
	errs := db.QueryRow(query, param...).Scan(&data.ID)
	if errs != nil && errs != sql.ErrNoRows {
		err = generateErrorModel(errs)
	}
	return
}
func GetParent(db *sql.DB, ncID int64, data *model.DataUserLevel, listData []interface{}) (err model.ErrorModel) {
	m := make(map[string]interface{})
	m["nexchief_account_id"] = ncID
	query, param := GetQueryParent("geo_tree", "parent", "code", m, listData, 0)

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
			case "parent":
				data.ParentID = id.Int64
			}
		}
	}
	return err
}

