package dao

import (
	"database/sql"
	"rohmat.co.id/model"
	"time"
)

func updateNexsellerChannelHierarchy(tx *sql.Tx, name, schema string, data *model.NexsellerCustomer) (err model.ErrorModel) {
	if data.NexsellerChannelHierarchyID == 0 {
		return
	}
	query := "UPDATE " + getSchema("nexseller_channel_hierarchy", schema) +
		" SET name = $1, updated_at = $2 WHERE id = $3 "
	_, errs := tx.Exec(query, name, time.Now(), data.NexsellerChannelHierarchyID)
	if errs != nil {
		err = generateErrorModel(errs)
		return
	}
	return
}
func condition1(tx *sql.Tx, ncID, mnID int64, schema string, data *model.NexsellerCustomer) (err model.ErrorModel) {
	var (
		query    string
		param    []interface{}
		tempKey  sql.NullString
		tempID   sql.NullInt64
		tempName sql.NullString
	)
	m := make(map[string]interface{})
	m["nexchief_account_id"] = ncID
	listData := []interface{}{data.MarketSegment}
	if data.TypeCode != "" {
		listData = append(listData, data.TypeCode)
	}
	if data.SubTypeCode != "" {
		listData = append(listData, data.SubTypeCode)
	}
	query, param = GetQueryParent("channel_hierarchy", "ch", "code", "name", m, listData, len(param))
	errs := tx.QueryRow(query, param...).Scan(&tempKey, &tempID, &tempName)
	if errs != nil && errs != sql.ErrNoRows {
		err = generateErrorModel(errs)
		return
	}

	if tempID.Int64 == 0 {
		return
	}
	err = updateNexsellerChannelHierarchy(tx, tempName.String, schema, data)
	return
}

func condition2(tx *sql.Tx, ncID, mnID int64, schema string, data *model.NexsellerCustomer) (err model.ErrorModel) {
	var (
		tempID   sql.NullInt64
		tempName sql.NullString
	)

	query := "SELECT id, name FROM " + getSchema("nexseller_channel_hierarchy", schema) +
		"WHERE mapping_nexseller_id = $1 AND code = $2 AND level = 2"
	errs := tx.QueryRow(query, mnID, data.TypeCode).Scan(&tempID, &tempName)
	if errs != nil && errs != sql.ErrNoRows {
		err = generateErrorModel(errs)
		return
	}
	data.NexsellerChannelHierarchyID = tempID.Int64
	if tempName.String != "" {
		return
	}

	query = "SELECT id, name FROM channel_hierarchy WHERE nexchief_account_id = $1 AND code = $2 AND level = 2"
	errs = tx.QueryRow(query, ncID, data.TypeCode).Scan(&tempID, &tempName)
	if errs != nil && errs != sql.ErrNoRows {
		err = generateErrorModel(errs)
		return
	}
	if tempID.Int64 == 0 {
		return
	}
	err = updateNexsellerChannelHierarchy(tx, tempName.String, schema, data)
	return
}

func condition3(tx *sql.Tx, ncID, mnID int64, schema string, data *model.NexsellerCustomer) (err model.ErrorModel) {
	var (
		tempID   sql.NullInt64
		tempName sql.NullString
	)

	query := "SELECT id, name FROM " + getSchema("nexseller_channel_hierarchy", schema) +
		"WHERE mapping_nexseller_id = $1 AND code = $2 AND level = 3"
	errs := tx.QueryRow(query, mnID, data.SubTypeCode).Scan(&tempID, &tempName)
	if errs != nil && errs != sql.ErrNoRows {
		err = generateErrorModel(errs)
		return
	}
	data.NexsellerChannelHierarchyID = tempID.Int64
	if tempName.String != "" {
		return
	}

	query = "SELECT id, name FROM channel_hierarchy WHERE nexchief_account_id = $1 AND code = $2 AND level = 3"
	errs = tx.QueryRow(query, ncID, data.SubTypeCode).Scan(&tempID, &tempName)
	if errs != nil && errs != sql.ErrNoRows {
		err = generateErrorModel(errs)
		return
	}
	if tempID.Int64 == 0 {
		return
	}
	err = updateNexsellerChannelHierarchy(tx, tempName.String, schema, data)
	return
}

func GetNexsellerChannelHierarchy(tx *sql.Tx, ncID, mnID int64, schema string, data *model.NexsellerCustomer) (err model.ErrorModel) {
	if data.MarketSegment != "" {
		err = condition1(tx, ncID, mnID, schema, data)
	} else if data.MarketSegment == "" && data.TypeCode != "" && data.SubTypeCode == "" {
		err = condition2(tx, ncID, mnID, schema, data)
	} else if data.MarketSegment == "" && data.TypeCode == "" && data.SubTypeCode != "" {
		err = condition3(tx, ncID, mnID, schema, data)
	} else if data.MarketSegment == "" && data.TypeCode != "" && data.SubTypeCode != "" {
		err = condition3(tx, ncID, mnID, schema, data)
	}
	return
}
