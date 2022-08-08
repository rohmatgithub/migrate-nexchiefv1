package dao

import (
	"database/sql"
	"rohmat.co.id/model"
)

func GetMappingNexseller(db *sql.DB, ncId int64, data *model.MappingNexseller) (err model.ErrorModel) {
	query :=
		"SELECT 'mn', id FROM mapping_nexseller where nexchief_account_id = $1 AND code = $2 " +
			"UNION ALL " +
			"SELECT 'pg', id FROM price_group where nexchief_account_id = $1 AND code = $3 " +
			"UNION ALL " +
			"SELECT 'pc', id FROM product_category where nexchief_account_id = $1 AND code = $4 " +
			"UNION ALL "
	m := make(map[string]interface{})
	m["nexchief_account_id"] = ncId
	tempQuery, tempParam := GetQueryParent("geo_tree", "gt", "code", "", m, []interface{}{
		data.UserLevel1, data.UserLevel2, data.UserLevel3, data.UserLevel4, data.UserLevel5,
	}, 4)

	query += tempQuery
	param := []interface{}{ncId, data.Code, data.PriceGroupCode, data.ProductCategoryCode}
	param = append(param, tempParam...)

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
			case "mn":
				data.ID = id.Int64
			case "pg":
				data.PriceGroupID = id.Int64
			case "pc":
				data.ProductCategoryID = id.Int64
			case "gt":
				data.GeoTreeID = id.Int64
			}
		}
	}
	return
}

func InsertMappingNexseller(db *sql.DB, ncID int64, data *model.MappingNexseller) (err model.ErrorModel) {
	query :=
		"INSERT INTO mapping_nexseller " +
			"(nexchief_account_id, code, geo_tree_id, " +
			"is_product_mapping, is_salesman_mapping, is_customer_mapping," +
			"hosting_only, is_generate_delivery, email_data, sync_method, " +
			"additional_info, active_date, resign_date, " +
			"price_group_id, product_category_id, nexseller_product_class_from," +
			"nexseller_product_class_thru, last_dms_sync, last_sfa_sync, " +
			"gromart_merchant_id, prefix_deleted, nd6_closed_date, " +
			"socket_user_id, socket_password, socket_status, " +
			"email_to, email_to_cc, product_registration_date, product_valid_thru )" +
			"VALUES " +
			"($1, $2, $3, " +
			"$4, $5, $6, " +
			"$7, $8, $9, $10, " +
			"$11, $12, $13, " +
			"$14, $15, $16, " +
			"$17, $18, $19, " +
			"$20, $21, $22, " +
			"$23, $24, $25, " +
			"$26, $27, $28, $29)"
	param := []interface{}{
		ncID, data.Code, data.GeoTreeID,
		getBool(data.IsProductMapping), getBool(data.IsSalesmanMapping), getBool(data.IsCustomerMapping),
		getBool(data.HostingOnly), getBool(data.IsGenerateDelivery), data.EmailData, data.SyncMethod,
		StructToJSON(model.AddInfo{
			MappingField1:  data.MappingField1,
			MappingField2:  data.MappingField2,
			MappingField3:  data.MappingField3,
			MappingField4:  data.MappingField4,
			MappingField5:  data.MappingField5,
			MappingField6:  data.MappingField6,
			MappingField7:  data.MappingField7,
			MappingField8:  data.MappingField8,
			MappingField9:  data.MappingField9,
			MappingField10: data.MappingField10,
		}), getNull(data.ActiveDate), getNull(data.ResignDate),
		data.PriceGroupID, data.ProductCategoryID, data.ProductClassFrom,
		data.ProductClassThru, getNull(data.LastDMSSyncStr), getNull(data.LastSFASyncStr),
		data.GromartMerchantID, data.PrefixDeleted, getNull(data.Nd6ClosedDate),
		data.SocketUserID, data.SocketPassword, getNull(data.SocketStatus),
		data.EmailTo, data.EmailToCC, getNull(data.ProductRegistrationDate), getNull(data.ProductValidThru),
	}
	_, errS := db.Exec(query, param...)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return err
}
