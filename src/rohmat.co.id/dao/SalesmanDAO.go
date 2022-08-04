package dao

import (
	"database/sql"
	"rohmat.co.id/model"
)

func GetSalesmanFKID(db *sql.DB, mnID int64, data *model.Salesman, schema string) (err model.ErrorModel) {
	query := "SELECT 'sl', id from " + getSchema("salesman", schema) + " WHERE mapping_nexseller_id = $1 AND code = $2 " +
		"UNION ALL " +
		"SELECT 'st', id FROM " + getSchema("principal_salesman_type", schema) + " WHERE code = $3 "
	param := []interface{}{mnID, data.Code, data.PrincipalSalesmanType}
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
			case "sl":
				data.ID = id.Int64
			case "st":
				data.PrincipalSalesmanTypeID = id.Int64
			}
		}
	}
	return
}

func GetPersonProfileID(db *sql.DB, data *model.Salesman) (result int64, err model.ErrorModel) {
	query := "SELECT id from person_profile WHERE " +
		"first_name = $1 AND nik = $2 AND address_1 = $3 " +
		"AND address_2 = $4 AND district_name = $5 AND phone = $6 " +
		"AND email = $7 "
	param := []interface{}{data.Name, data.Nik, data.Address1,
		data.Address2, data.City, data.Phone,
		data.Email}
	errS := db.QueryRow(query, param...).Scan(&result)
	if errS != nil && errS != sql.ErrNoRows {
		err = generateErrorModel(errS)
	}
	return
}

func InsertPersonProfile(db *sql.DB, data *model.Salesman) (result int64, err model.ErrorModel) {
	query := "INSERT INTO person_profile " +
		"(first_name, nik, address_1, " +
		"address_2, district_name, phone, " +
		"email) " +
		"VALUES " +
		"($1, $2, $3, " +
		"$4, $5, $6, " +
		"$7) " +
		"RETURNING id"
	param := []interface{}{
		data.Name, data.Nik, data.Address1,
		data.Address2, data.City, data.Phone,
		data.Email,
	}
	errS := db.QueryRow(query, param...).Scan(&result)
	if errS != nil {
		err = generateErrorModel(errS)
		return
	}
	return
}

func InsertSalesman(db *sql.DB, data *model.Salesman, nc model.NexchiefAccount, personProfileID, mnID int64) (err model.ErrorModel) {
	query := "INSERT INTO " + getSchema("Salesman", nc.Schema.String) +
		" (nexchief_account_id, mapping_nexseller_id, code, " +
		"name, address_1, address_2, " +
		"city, phone, email, " +
		"nik, join_date, resign_date, " +
		"type, category, salesman_nexmile_on," +
		"last_sync_dms, imei_number, imei_number_2," +
		"registration_date, registration_city, registration_latitude, " +
		"registration_longitude, salesman_group, principal_salesman_type_id," +
		"nexmile_version, nexmile_valid_thru, nexmile_device_id, " +
		"status, person_profile_id, pk_checksum) " +
		"VALUES " +
		"($1, $2, $3," +
		"$4, $5, $6, " +
		"$7, $8, $9, " +
		"$10, $11, $12, " +
		"$13, $14, $15, " +
		"$16, $17, $18, " +
		"$19, $20, $21, " +
		"$22, $23, $24, " +
		"$25, $26, $27, " +
		"$28, $29, $30) "
	param := []interface{}{
		nc.ID.Int64, mnID, data.Code,
		data.Name, data.Address1, data.Address2,
		data.City, data.Phone, data.Email,
		data.Nik, getNull(data.JoinDate), getNull(data.ResignDate),
		data.Type, data.Category, data.NexileOn,
		getNull(data.LastSync), data.ImeiNumber, data.ImeiNumber2,
		getNull(data.RegistrationDate), data.RegistrationCity, data.RegisLat,
		data.RegisLong, data.Group, data.PrincipalSalesmanTypeID,
		data.NexmileVersion, getNull(data.NexmileValidThru), data.NexmileDeviceID,
		data.Status, personProfileID, data.PkChecksum,
	}
	_, errS := db.Exec(query, param...)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return
}
