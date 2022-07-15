package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"rohmat.co.id/model"
)

func GetMappingNexsellerID(db *sql.DB, ncID int64, code string) (result int64, err model.ErrorModel) {
	query := "SELECT id from mapping_nexseller WHERE " +
		"nexchief_account_id = $1 AND code = $2"
	param := []interface{}{ncID, code}
	errS := db.QueryRow(query, param...).Scan(&result)
	if errS != nil && errS != sql.ErrNoRows {
		err.Error = errS
		err.Code = 500
		return
	}else if result == 0 {
		err.Error = errors.New(fmt.Sprintf("mapping nexseller : %s not exists", code))
		return
	}
	return
}

func GetNexProductFK(db *sql.DB, schema string, mnID int64, data *model.NexsellerProduct) (err model.ErrorModel) {
	query :=
		"SELECT 'np', id FROM " + getSchema("nexseller_product", schema) + " where mapping_nexseller_id = $1 AND product_code = $2 " +
			"UNION ALL " +
			"SELECT 'nv', id FROM " + getSchema("nexseller_vendor", schema) + " where mapping_nexseller_id = $1 AND code = $3 " +
			"UNION ALL " +
			"SELECT 'nd', id FROM " + getSchema("nexseller_division", schema) + " where mapping_nexseller_id = $1 AND code = $4 " +
			"UNION ALL " +
			"SELECT 'pp', id FROM " + getSchema("product", schema) + " where principal_product_code = $5 "
	param := []interface{}{mnID, data.Code, data.VendorCode, data.DivisionCode, data.PrincipalProductCode}

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
			case "np":
				data.ID = id.Int64
			case "nv":
				data.VendorID = id.Int64
			case "nd":
				data.DivisionID = id.Int64
			case "pp":
				data.PrincipalProductID = id.Int64
			}
		}
	}
	return
}

func InsertNexsellerProduct(db *sql.DB, nc model.NexchiefAccount, mnID int64, model *model.NexsellerProduct) (err model.ErrorModel) {
	query := "INSERT INTO " + getSchema("nexseller_product", nc.Schema.String) +
		"(product_code, product_id, nexchief_account_id, " +
		"mapping_nexseller_id, name, packaging, " +
		"uom_1, uom_2, uom_3, " +
		"uom_4, conversion_1_to_4, conversion_2_to_4, " +
		"conversion_3_to_4, status, buying_price, " +
		"selling_price, nexchief_ratio, nexseller_vendor_id, " +
		"nexseller_division_id) " +
		"VALUES " +
		"($1, $2, $3, " +
		"$4, $5, $6, " +
		"$7, $8, $9," +
		"$10, $11, $12," +
		"$13, $14, $15," +
		"$16, $17, $18," +
		"$19) "
	param := []interface{}{
		model.Code, model.PrincipalProductID, nc.ID.Int64,
		mnID, model.Name, model.Packaging,
		model.Uom1, model.Uom2, model.Uom3,
		model.Uom4, model.Conversion1to4, model.Conversion2to4,
		model.Conversion3to4, model.Status, model.BuyingPrice,
		model.SellingPrice, model.NexchiefRatio, model.VendorID,
		model.DivisionID,
	}
	_, errS := db.Exec(query, param...)
	if errS != nil {
		err = generateErrorModel(errS)
	}
	return
}
