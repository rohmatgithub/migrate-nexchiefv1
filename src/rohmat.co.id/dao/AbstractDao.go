package dao

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"rohmat.co.id/model"
	"strconv"
	"time"
)

var MapNexchiefAccount = make(map[string]model.NexchiefAccount)

func GetQueryParent(tableName, key, fieldClause, addSelect string, clauseMustCheck map[string]interface{}, listData []interface{}, lengthParam int) (query string, queryParam []interface{}) {
	query += fmt.Sprintf("SELECT '%s', ", key)
	clause := " "
	tempLengthParam := lengthParam
	l := 0
	for i := len(listData); i > 0; i-- {
		if listData[i-1] == "" && lengthParam == tempLengthParam {
			continue
		}
		l++
		if l == 1 {
			query += fmt.Sprintf(" lv%d.id ", i)
			if addSelect != "" {
				query += fmt.Sprintf(", lv%d.%s ", i, addSelect)
			}

			query += fmt.Sprintf(" FROM %s lv%d", tableName, i)
		} else {
			query += fmt.Sprintf(" LEFT JOIN %s lv%d ON lv%d.id = lv%d.parent_id ", tableName, i, i, i+1)
		}
		lengthParam++
		clause += fmt.Sprintf(" lv%d.%s = $%d ", i, fieldClause, lengthParam)
		if i > 1 {
			clause += "AND"
		}
		queryParam = append(queryParam, listData[i-1])
	}
	lengthParam++
	query += " WHERE " + clause + " AND lv" + strconv.Itoa(l) + ".level = $" + strconv.Itoa(lengthParam)
	queryParam = append(queryParam, l)
	for s, i := range clauseMustCheck {
		lengthParam++
		query += fmt.Sprintf(" AND lv%d.%s = $%d ", l, s, lengthParam)
		queryParam = append(queryParam, i)
	}
	return
}

func GetNexchiefAccountID(db *sql.DB, code string) (result model.NexchiefAccount, errorModel model.ErrorModel) {
	result = MapNexchiefAccount[code]

	// get to db if null in map
	if result.ID.Int64 == 0 {
		fmt.Println("get nexchief account from db :", code)
		query :=
			" SELECT id, schema from nexchief_account WHERE code = $1 "
		param := []interface{}{code}

		err := db.QueryRow(query, param...).Scan(&result.ID, &result.Schema)
		if err != nil && err != sql.ErrNoRows {
			generateErrorModel(err)
			return
		}
		if result.ID.Int64 > 0 && result.Schema.String != "" {
			MapNexchiefAccount[code] = result
		} else {
			errorModel.Error = errors.New("nexchief account not found")
			errorModel.Code = 400
		}
	}
	return
}

func getSchema(tableName, schema string) string {
	if schema != "" {
		return schema + "." + tableName
	}
	return tableName
}

func TimeToString(time time.Time) string {
	if time.IsZero() {
		return ""
	}
	return time.Format("2006-01-02T15:04:05.000000Z")
}

func getNull(str string) interface{} {
	if str == "" {
		return nil
	}
	return str
}
func getBool(str string) bool {
	if str == "Y" {
		return true
	}
	return false
}

func generateErrorModel(err error) model.ErrorModel {
	return model.ErrorModel{
		Error: err,
		Code:  500,
	}
}

func StructToJSON(input interface{}) (output string) {
	b, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
		output = ""
		return
	}
	output = string(b)
	return
}
