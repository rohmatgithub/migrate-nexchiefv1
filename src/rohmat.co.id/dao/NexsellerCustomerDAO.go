package dao

import (
	"database/sql"
	"rohmat.co.id/model"
	"time"
)

func GetNexsellerCustomerFK(db *sql.DB, nc model.NexchiefAccount, data *model.NexsellerCustomer, mnID int64) (errModel model.ErrorModel) {
	query := "SELECT 'nc', id, name FROM " + getSchema("nexseller_customer", nc.Schema.String) + " WHERE mapping_nexseller_id = $1 AND code = $2 " +
		" UNION ALL " +
		"SELECT 'cg', id, customer_group_name FROM " + getSchema("nexseller_customer_group", nc.Schema.String) + " WHERE mapping_nexseller_id = $1 AND customer_group_id = $3 " +
		" UNION ALL " +
		"SELECT 'cc', id, name FROM " + getSchema("nexseller_customer_category", nc.Schema.String) + " WHERE mapping_nexseller_id = $1 AND code = $4 " +
		" UNION ALL " +
		"SELECT 'np', id, name FROM " + getSchema("nexseller_province", nc.Schema.String) + " WHERE mapping_nexseller_id = $1 AND code = $5 " +
		" UNION ALL " +
		"SELECT 'nd', id, name FROM " + getSchema("nexseller_district", nc.Schema.String) + " WHERE mapping_nexseller_id = $1 AND code = $6 " +
		" UNION ALL " +
		"SELECT 'nsd', id, name FROM " + getSchema("nexseller_sub_district", nc.Schema.String) + " WHERE mapping_nexseller_id = $1 AND code = $7 " +
		" UNION ALL " +
		"SELECT 'nuv', id, name FROM " + getSchema("nexseller_urban_village", nc.Schema.String) + " WHERE mapping_nexseller_id = $1 AND code = $8 " +
		" UNION ALL " +
		"SELECT 'nsl', id, name FROM " + getSchema("nexseller_store_location", nc.Schema.String) + " WHERE mapping_nexseller_id = $1 AND code = $9 " +
		" UNION ALL " +
		"SELECT 'nss', id, name FROM " + getSchema("nexseller_store_status", nc.Schema.String) + " WHERE mapping_nexseller_id = $1 AND code = $10 " +
		" UNION ALL " +
		"SELECT 'ni', id, name FROM " + getSchema("nexseller_island", nc.Schema.String) + " WHERE mapping_nexseller_id = $1 AND code = $11 " +
		" UNION ALL " +
		"SELECT 'nuc1', id, name FROM " + getSchema("nexseller_user_category", nc.Schema.String) + " WHERE mapping_nexseller_id = $1 AND code = $12 AND type = 1 " +
		" UNION ALL " +
		"SELECT 'nuc2', id, name FROM " + getSchema("nexseller_user_category", nc.Schema.String) + " WHERE mapping_nexseller_id = $1 AND code = $13 AND type = 2 " +
		" UNION ALL " +
		"SELECT 'nuc3', id, name FROM " + getSchema("nexseller_user_category", nc.Schema.String) + " WHERE mapping_nexseller_id = $1 AND code = $14 AND type = 3 "
	param := []interface{}{
		mnID, data.Code, data.GroupCode, data.CategoryCode, data.ProvinceCode, data.DistrictCode,
		data.SubDistrictCode, data.UrbanVillageCode, data.StoreLocationCode, data.StoreStatusCode,
		data.IslandCode, data.UserCategory1Code, data.UserCategory2Code, data.UserCategory3Code,
	}

	m := make(map[string]interface{})
	m["mapping_nexseller_id"] = mnID
	// query sub area hierarchy
	if data.AreaCode != "" {
		query += " UNION ALL "
		listData := []interface{}{data.AreaCode}
		if data.SubAreaCode != "" {
			listData = append(listData, data.SubAreaCode)
		}
		tempQuery, tempParam := GetQueryParent(getSchema("nexseller_area_hierarchy", nc.Schema.String), "nah", "code", "name", m, listData, len(param))
		query += tempQuery
		param = append(param, tempParam...)
	}
	if data.MarketSegment != "" {
		query += " UNION ALL "
		listData := []interface{}{data.MarketSegment}
		if data.TypeCode != "" {
			listData = append(listData, data.TypeCode)
		}
		if data.SubTypeCode != "" {
			listData = append(listData, data.SubTypeCode)
		}
		tempQuery, tempParam := GetQueryParent(getSchema("nexseller_channel_hierarchy", nc.Schema.String), "nch", "code", "name", m, listData, len(param))
		query += tempQuery
		param = append(param, tempParam...)
	}

	rows, err := db.Query(query, param...)
	if err != nil {
		return generateErrorModel(err)
	}
	if rows != nil {
		defer func() {
			err = rows.Close()
			if err != nil {
				errModel = generateErrorModel(err)
				return
			}
		}()

		for rows.Next() {
			var (
				key  string
				id   sql.NullInt64
				name sql.NullString
			)

			err = rows.Scan(&key, &id, &name)
			if err != nil {
				return generateErrorModel(err)
			}
			switch key {
			case "nc":
				data.ID = id.Int64
			case "cg":
				data.GroupID = id.Int64
				data.GroupName = name.String
			case "cc":
				data.CategoryID = id.Int64
				data.CategoryName = name.String
			case "np":
				data.ProvinceID = id.Int64
				data.ProvinceName = name.String
			case "nd":
				data.DistrictID = id.Int64
				data.DistrictName = name.String
			case "nsd":
				data.SubDistrictID = id.Int64
				data.SubDistrictName = name.String
			case "nuv":
				data.UrbanVillageID = id.Int64
				data.UrbanVillageName = name.String
			case "nsl":
				data.StoreLocationID = id.Int64
				data.StoreLocationName = name.String
			case "nss":
				data.StoreStatusID = id.Int64
				data.StoreStaturName = name.String
			case "nah":
				data.NexsellerAreaHierarchyID = id.Int64
			case "nch":
				data.NexsellerChannelHierarchyID = id.Int64
				data.NexsellerChanneHierarchyName = name.String
			case "ni":
				data.IslandID = id.Int64
				data.IslandName = name.String
			case "nuc1":
				data.UserCategory1ID = id.Int64
				data.UserCategory1Name = name.String
			case "nuc2":
				data.UserCategory2ID = id.Int64
				data.UserCategory2Name = name.String
			case "nuc3":
				data.UserCategory3ID = id.Int64
				data.UserCategory3Name = name.String
			}
		}
	}
	return
}

func GetCompanyProfile(db *sql.Tx, data *model.NexsellerCustomer) (err model.ErrorModel) {
	query := "SELECT id FROM company_profile " +
		"WHERE name = $1 AND address_1 = $2 AND address_2 = $3 " +
		"AND address_3 = $4 AND district = $5 AND phone = $6 " +
		"AND fax = $7 AND email = $8"
	param := []interface{}{
		data.Name, data.Address1, data.Address2,
		data.Address3, data.DistrictCode, data.Phone,
		data.Fax, data.Email,
	}
	errS := db.QueryRow(query, param...).Scan(&data.CompanyProfileID)
	if errS != nil && errS != sql.ErrNoRows {
		err = generateErrorModel(errS)
		return err
	} else if data.CompanyProfileID == 0 {
		err = insertCompanyProfileFromCustomer(db, data)
	}
	return
}

func insertCompanyProfileFromCustomer(db *sql.Tx, data *model.NexsellerCustomer) (err model.ErrorModel) {
	query := "INSERT INTO company_profile " +
		"( name, address_1, address_2, " +
		"address_3, district, phone, " +
		"fax, email) " +
		"VALUES " +
		"($1, $2, $3, " +
		"$4, $5, $6, " +
		"$7, $8 ) RETURNING id"
	param := []interface{}{
		data.Name, data.Address1, data.Address2,
		data.Address3, data.DistrictCode, data.Phone,
		data.Fax, data.Email,
	}
	errS := db.QueryRow(query, param...).Scan(&data.CompanyProfileID)
	if errS != nil && errS != sql.ErrNoRows {
		err = generateErrorModel(errS)
		return err
	}
	return
}
func InsertNexsellerCustomer(db *sql.Tx, nc model.NexchiefAccount, mnID int64, data *model.NexsellerCustomer) (err model.ErrorModel) {
	query := "INSERT INTO " + getSchema("nexseller_customer", nc.Schema.String) +
		" (nexchief_account_id, mapping_nexseller_id, company_profile_id," +
		"code, name, address_1, " +
		"address_2, address_3, city, " +
		"phone, msg_number, fax," +
		"email, nexseller_area_hierarchy_id, nexseller_channel_hierarchy_id," +
		"nexseller_customer_category_id, nexseller_customer_group_id, class, " +
		"status, is_bumn, is_pkp, " +
		"latitude, longitude, additional_info," +
		"island_id, province_id, district_id, " +
		"sub_district_id, urban_village_id, join_date, " +
		"last_sync_dms, store_location_id, store_status_id, " +
		"location_remark, flag_verified, user_category_id_1, " +
		"user_category_id_2, user_category_id_3, pk_checksum ) " +
		"VALUES " +
		"($1, $2, $3, " +
		"$4, $5, $6, " +
		"$7, $8, $9, " +
		"$10, $11, $12," +
		"$13, $14, $15," +
		"$16, $17, $18," +
		"$19, $20, $21, " +
		"$22, $23, $24, " +
		"$25, $26, $27, " +
		"$28, $29, $30," +
		"$31, $32, $33," +
		"$34, $35, $36," +
		"$37, $38, $39) "
	param := []interface{}{
		nc.ID.Int64, mnID, data.CompanyProfileID,
		data.Code, data.Name, data.Address1,
		data.Address2, data.Address3, data.City,
		data.Phone, data.MsgNumber, data.Fax,
		data.Email, data.NexsellerAreaHierarchyID, data.NexsellerChannelHierarchyID,
		data.CategoryID, data.GroupID, data.Class,
		data.Status, data.IsBUMN, data.IsPKP,
		data.Latitude, data.Longitude, convertToAdditionalInfo(data),
		data.IslandID, data.ProvinceID, data.DistrictID,
		data.SubDistrictID, data.UrbanVillageID, getNull(data.JoinDate),
		getNull(data.LastSync), data.StoreLocationID, data.StoreStatusID,
		data.LocationRemark, data.FlagVerified, data.UserCategory1ID,
		data.UserCategory2ID, data.UserCategory3ID, data.PkChecksum,
	}
	stmt, errS := db.Prepare(query)
	if errS != nil {
		err = generateErrorModel(errS)
		return err
	}
	defer stmt.Close()
	_, errS = stmt.Exec(param...)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return
}

func UpdateNexsellerCustomer(db *sql.Tx, nc model.NexchiefAccount, mnID int64, data *model.NexsellerCustomer) (err model.ErrorModel) {
	query := "UPDATE " + getSchema("nexseller_customer", nc.Schema.String) +
		" SET " +
		"nexchief_account_id = $1, mapping_nexseller_id = $2, company_profile_id = $3," +
		"code = $4, name = $5, address_1 = $6, " +
		"address_2 = $7, address_3 = $8, city = $9, " +
		"phone = $10, msg_number = $11, fax = $12," +
		"email = $13, nexseller_area_hierarchy_id = $14, nexseller_channel_hierarchy_id = $15," +
		"nexseller_customer_category_id = $16, nexseller_customer_group_id = $17, class = $18, " +
		"status = $19, is_bumn = $20, is_pkp = $21, " +
		"latitude = $22, longitude = $23, additional_info = $24," +
		"island_id = $25, province_id = $26, district_id = $27, " +
		"sub_district_id = $28, urban_village_id = $29, join_date = $30, " +
		"last_sync_dms = $31, store_location_id = $32, store_status_id = $33, " +
		"location_remark = $34, flag_verified = $35, user_category_id_1 = $36, " +
		"user_category_id_2 = $37, user_category_id_3 = $38, pk_checksum  = $39 " +
		"WHERE id = $40"
	param := []interface{}{
		nc.ID.Int64, mnID, data.CompanyProfileID,
		data.Code, data.Name, data.Address1,
		data.Address2, data.Address3, data.City,
		data.Phone, data.MsgNumber, data.Fax,
		data.Email, data.NexsellerAreaHierarchyID, data.NexsellerChannelHierarchyID,
		data.CategoryID, data.GroupID, data.Class,
		data.Status, data.IsBUMN, data.IsPKP,
		data.Latitude, data.Longitude, convertToAdditionalInfo(data),
		data.IslandID, data.ProvinceID, data.DistrictID,
		data.SubDistrictID, data.UrbanVillageID, getNull(data.JoinDate),
		getNull(data.LastSync), data.StoreLocationID, data.StoreStatusID,
		data.LocationRemark, data.FlagVerified, data.UserCategory1ID,
		data.UserCategory2ID, data.UserCategory3ID, data.PkChecksum,
		data.ID,
	}
	stmt, errS := db.Prepare(query)
	if errS != nil {
		err = generateErrorModel(errS)
		return err
	}
	defer stmt.Close()
	_, errS = stmt.Exec(param...)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return
}

func convertToAdditionalInfo(data *model.NexsellerCustomer) string {
	type AddInfoNexsellerCustomer struct {
		GromartFirstTransaction string `json:"gromart_first_transaction"` // format time 2006-01-02T15:04:05Z -> constanta.DefaultDateTimeFormat
		GromartLastTransaction  string `json:"gromart_last_transaction"`  // format time 2006-01-02T15:04:05Z -> constanta.DefaultDateTimeFormat
		IsCustomerTrade         string `json:"is_customer_trade"`
		IsPicos                 string `json:"is_picos"`
		IsTdWeb                 string `json:"is_td_web"`
		PicosDate               string `json:"picos_date"`
	}
	layout := "2006-01-02 15:04:05"
	gromartFirstTransaction, _ := time.Parse(layout, data.GromartFirstTransaction)
	gromartLastTransaction, _ := time.Parse(layout, data.GromartLastTransaction)
	result := AddInfoNexsellerCustomer{
		GromartFirstTransaction: TimeToString(gromartFirstTransaction),
		GromartLastTransaction:  TimeToString(gromartLastTransaction),
		IsCustomerTrade:         data.IsCustomerTrade,
		IsPicos:                 data.IsPicos,
		IsTdWeb:                 data.IsTdWeb,
		PicosDate:               data.PcosDate,
	}
	return StructToJSON(result)
}
