package dao

import (
	"database/sql"
	"rohmat.co.id/model"
)

func GetDistributor(db *sql.DB, data *model.Distributor) (result []model.MappingDistributor, err model.ErrorModel) {
	query :=
		"SELECT id, nexchief_account_id " +
			"FROM mapping_nexseller WHERE code = $1 "
	param := []interface{}{data.Code}
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
			var temp model.MappingDistributor

			errS = rows.Scan(&temp.ID, &temp.NcID)
			if errS != nil {
				err = generateErrorModel(errS)
				return
			}
			result = append(result, temp)
		}
	}
	return
}

func GetParentID(db *sql.DB, model *model.MappingDistributor) (id int64, err model.ErrorModel) {
	query :=
		"SELECT id FROM mapping_nexseller WHERE id = $1 AND nexchief_account_id = $2"
	param := []interface{}{model.ID, model.NcID}
	errS := db.QueryRow(query, param...).Scan(&id)
	if errS != nil && errS != sql.ErrNoRows {
		err = generateErrorModel(errS)
	}
	return
}

func GetIDCompanyProfile(db *sql.DB, model *model.Distributor) (id int64, err model.ErrorModel) {
	query :=
		"SELECT id FROM company_profile WHERE name = $1 AND address_1 = $2 AND address_2 = $3 " +
			"AND npwp = $4 AND district = $5 AND phone = $6 AND email = $7 "
	param := []interface{}{model.Name, model.Address1, model.Address2,
		model.Npwp, model.City, model.Phone, model.Email}
	errS := db.QueryRow(query, param...).Scan(&id)
	if errS != nil && errS != sql.ErrNoRows {
		err = generateErrorModel(errS)
	}
	return
}

func InsertCompanyProfile(db *sql.DB, userParam *model.Distributor) (id int64, err model.ErrorModel) {
	//funcName := "InsertDataGroup"
	query :=
		" INSERT INTO  company_profile " +
			"  (name, address_1, address_2," +
			"  district, email, phone, " +
			"  fax, npwp) " +
			" VALUES " +
			"($1, $2, $3, " +
			" $4, $5, $6, " +
			" $7, $8) " +
			"RETURNING id"
	param := []interface{}{
		userParam.Name, userParam.Address1, userParam.Address2,
		userParam.City, userParam.Email, userParam.Phone,
		userParam.Fax, userParam.Npwp}

	errS := db.QueryRow(query, param...).Scan(&id)
	if errS != nil && errS != sql.ErrNoRows {
		err = generateErrorModel(errS)
	}
	return
}

func UpdateMappingNexseller(db *sql.DB, id, cpID, parentID int64) (err model.ErrorModel) {
	query := "UPDATE mapping_nexseller SET " +
		"company_profile_id = $1, parent_id = $2 " +
		"WHERE id = $3 "
	_, errS := db.Exec(query, cpID, parentID, id)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return err
}

