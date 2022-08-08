package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"nexsoft.co.id/nexcommon/util"
	"rohmat.co.id/config"
	"rohmat.co.id/dao"
	"rohmat.co.id/model"
	"rohmat.co.id/serverconfig"
	"sync"
)

func StartSaveNexsellerCustomer(wg *sync.WaitGroup) {
	// 2021-01-01
	path := config.ApplicationConfiguration.GetDirPath().PathDir +
		config.ApplicationConfiguration.GetDirPath().Customer

	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, file := range files {
		if !file.IsDir() {
			wg.Add(1)
			go doStartWaitGroup(wg, path+"/"+file.Name(), fmt.Sprintf("nexseller customer %d", i))
		}
	}
}

func doStartWaitGroup(wg *sync.WaitGroup, path string, partData string) {
	defer wg.Done()

	StartReadFile(path, SaveNexsellerCustomer, partData)

}

func SaveNexsellerCustomer(db *sql.DB, dataByte []byte) (errorModel model.ErrorModel) {
	var (
		nexchiefAccount model.NexchiefAccount
		data            model.NexsellerCustomer
		mnID            int64
		tx              *sql.Tx
	)
	_ = json.Unmarshal(dataByte, &data)

	defer func() {
		if errorModel.Error != nil {
			_ = tx.Rollback()
		} else {
			errs := tx.Commit()
			if errs != nil {
				errorModel.Error = errs
				errorModel.Code = 500
			}
		}
	}()

	tx, errorModel.Error = db.Begin()
	if errorModel.Error != nil {
		return
	}
	nexchiefAccount, errorModel = dao.GetNexchiefAccountID(db, data.NcCode)
	if errorModel.Error != nil {
		return
	}
	// get mapping nexseller
	mnID, errorModel = dao.GetMappingNexsellerID(db, nexchiefAccount.ID.Int64, data.MnCode)
	if errorModel.Error != nil {
		return
	}
	// get data FK
	errorModel = dao.GetNexsellerCustomerFK(db, nexchiefAccount, &data, mnID)
	if errorModel.Error != nil {
		return
	}
	errorModel = validateFKCustomer(tx, nexchiefAccount, mnID, &data)
	if errorModel.Error != nil {
		return
	}

	pkString := data.NcCode + data.MnCode + data.Code
	data.PkChecksum = util.CheckSumWithXXHASH([]byte(pkString))
	if data.ID == 0 {
		// get company profile id
		errorModel = dao.GetCompanyProfile(tx, &data)
		if errorModel.Error != nil {
			return
		}
		errorModel = dao.InsertNexsellerCustomer(tx, nexchiefAccount, mnID, &data)
		if errorModel.Error != nil {
			return
		}
	} else {
		errorModel = dao.UpdateNexsellerCustomer(tx, nexchiefAccount, mnID, &data)
	}
	return
}

// FK nexseller customer
//- nexseller_island
//- nexseller_province
//- nexseller_district
//- nexseller_sub_district
//- nexseller_urban_village
// - nexseller_area_hierarchy_id
//- nexseller_channel_hierarchy_id
//- nexseller_customer_category_id
//- nexseller_customer_group_id
//- nexseller_customer_key_account_id

func validateFKCustomer(tx *sql.Tx, nc model.NexchiefAccount, mnID int64, data *model.NexsellerCustomer) (err model.ErrorModel) {
	db := serverconfig.ServerAttribute.DBConnection
	var masterModel model.MasterModel
	if data.IslandCode != "" && data.IslandName != "" {
		// get master island
		masterModel, err = dao.GetIsland(db, nc.ID.Int64, data.IslandCode)
		if err.Error != nil {
			return
		}
		// insert island
		masterModel.Code.String = data.IslandCode
		if data.IslandID == 0 {
			data.IslandID, err = dao.InsertNexsellerIsland(tx, nc.ID.Int64, mnID, nc.Schema.String, &masterModel)
		} else {
			err = dao.UpdateNexsellerIsland(tx, data.IslandID, nc.Schema.String, &masterModel)
		}
		if err.Error != nil {
			return
		}
	}

	if data.UrbanVillageCode != "" && data.SubDistrictCode == "" {
		// get master sub district by urban village
		data.SubDistrictCode, err = dao.GetSubDistrictByUrbanVillage(db, data.UrbanVillageCode)
		if err.Error != nil {
			return
		}
		// get nexseller sub district
		data.SubDistrictID, data.SubDistrictName, err = dao.GetNexsellerSubDistrict(db, nc.Schema.String, data.SubDistrictCode, mnID)
		if err.Error != nil {
			return
		}
	}

	if data.SubDistrictCode != "" && data.DistrictCode == "" {
		// get district by sub district
		data.DistrictCode, err = dao.GetDistrictBySubDistrict(db, data.SubDistrictCode)
		if err.Error != nil {
			return
		}
		// get nexseller district
		data.DistrictID, data.DistrictName, err = dao.GetNexsellerDistrict(db, nc.Schema.String, data.DistrictCode, mnID)
		if err.Error != nil {
			return
		}
	}

	if data.DistrictCode != "" && data.ProvinceCode == "" {
		// get province by district
		data.ProvinceCode, err = dao.GetProvinceByDistrict(db, data.DistrictCode)
		if err.Error != nil {
			return
		}
		// get nexseller province
		data.ProvinceID, data.ProvinceName, err = dao.GetNexsellerProvince(db, nc.Schema.String, data.ProvinceCode, mnID)
		if err.Error != nil {
			return
		}
	}

	if data.ProvinceCode != "" && data.ProvinceName == "" {
		// get name dari master
		masterModel, err = dao.GetProvince(db, data.ProvinceCode)
		if err.Error != nil {
			return
		}
		if data.ProvinceID == 0 {
			// insert nexseller_province
			data.ProvinceID, err = dao.InsertNexsellerProvince(tx, nc.ID.Int64, mnID, nc.Schema.String, &masterModel)
		} else {
			// update nexseller province
			err = dao.UpdateNexsellerProvince(tx, data.ProvinceID, nc.Schema.String, &masterModel)
		}
		if err.Error != nil {
			return
		}
	}

	if data.DistrictCode != "" && data.DistrictName == "" {
		// get name dari master
		masterModel, err = dao.GetDistrict(db, data.DistrictCode)
		if err.Error != nil {
			return
		}
		if data.DistrictID == 0 {
			masterModel.ParentID.Int64 = data.ProvinceID
			// insert nexseller_district
			data.DistrictID, err = dao.InsertNexsellerDistrict(tx, nc.ID.Int64, mnID, nc.Schema.String, &masterModel)
		} else {
			// update nexseller_district
			err = dao.UpdateNexsellerDistrict(tx, data.DistrictID, nc.Schema.String, &masterModel)
		}
		if err.Error != nil {
			return
		}
	}

	if data.SubDistrictCode != "" && data.SubDistrictName == "" {
		// get name dari master
		masterModel, err = dao.GetSubDistrict(db, data.SubDistrictCode)
		if err.Error != nil {
			return
		}
		if data.SubDistrictID == 0 {
			// insert nexseller_sub_district
			data.SubDistrictID, err = dao.InsertNexsellerSubDistrict(tx, nc.ID.Int64, mnID, nc.Schema.String, &masterModel)
		} else {
			// update nexseller_sub_district
			err = dao.UpdateNexsellerSubDistrict(tx, data.SubDistrictID, nc.Schema.String, &masterModel)
		}
		if err.Error != nil {
			return
		}
	}

	if data.UrbanVillageCode != "" && data.UrbanVillageName == "" {
		// get name dari master
		masterModel, err = dao.GetUrbanVillage(db, data.UrbanVillageCode)
		if err.Error != nil {
			return
		}
		if data.UrbanVillageID == 0 {
			// insert
			data.UrbanVillageID, err = dao.InsertNexsellerUrbanVillage(tx, nc.ID.Int64, mnID, nc.Schema.String, &masterModel)
		} else {
			// update
			err = dao.UpdateNexsellerUrbanVillage(tx, data.UrbanVillageID, nc.Schema.String, &masterModel)
		}
		if err.Error != nil {
			return
		}
	}

	//- nexseller_channel_hierarchy_id
	if data.MarketSegment != "" && data.NexsellerChanneHierarchyName == "" {
		err = dao.GetNexsellerChannelHierarchy(tx, nc.ID.Int64, mnID, nc.Schema.String, data)
		if err.Error != nil {
			return
		}
	}
	//- nexseller_customer_category_id
	if data.CategoryCode != "" && data.CategoryName == "" {
		masterModel, err = dao.GetCustomerCategory(db, nc.ID.Int64, data)
		if err.Error != nil {
			return
		}
		if data.CategoryID == 0 {
			data.CategoryID, err = dao.InsertNexsellerCustomerCategory(tx, nc.ID.Int64, mnID, nc.Schema.String, &masterModel)
		} else {
			err = dao.UpdateNexsellerCustomerCategory(tx, data.CategoryID, nc.Schema.String, &masterModel)
		}
		if err.Error != nil {
			return
		}
	}
	//- nexseller_customer_group_id
	if data.GroupCode != "" && data.GroupName == "" {
		masterModel, err = dao.GetCustomerGroup(db, nc.ID.Int64, data)
		if err.Error != nil {
			return
		}
		if data.GroupID == 0 {
			data.GroupID, err = dao.InsertNexsellerCustomerGroup(tx, nc.ID.Int64, mnID, nc.Schema.String, &masterModel)
		} else {
			err = dao.UpdateNexsellerCustomerGroup(tx, data.GroupID, nc.Schema.String, &masterModel)
		}
		if err.Error != nil {
			return
		}
	}

	//- nexseller_user_category
	if data.UserCategory1Code != "" && data.UserCategory1Name == "" {
		masterModel, err = dao.GetUserCategory(db, nc.ID.Int64, 1, nc.Schema.String, data)
		if err.Error != nil {
			return
		}
		if data.UserCategory1ID == 0 {
			data.GroupID, err = dao.InsertNexsellerUserCategory(tx, nc.ID.Int64, mnID, 1, nc.Schema.String, &masterModel)
		} else {
			err = dao.UpdateNexsellerUserCategory(tx, data.UserCategory1ID, nc.Schema.String, &masterModel)
		}
		if err.Error != nil {
			return
		}
	}

	if data.UserCategory2Code != "" && data.UserCategory2Name == "" {
		masterModel, err = dao.GetUserCategory(db, nc.ID.Int64, 2, nc.Schema.String, data)
		if err.Error != nil {
			return
		}
		if data.UserCategory2ID == 0 {
			data.GroupID, err = dao.InsertNexsellerUserCategory(tx, nc.ID.Int64, mnID, 2, nc.Schema.String, &masterModel)
		} else {
			err = dao.UpdateNexsellerUserCategory(tx, data.UserCategory2ID, nc.Schema.String, &masterModel)
		}
		if err.Error != nil {
			return
		}
	}

	if data.UserCategory3Code != "" && data.UserCategory3Name == "" {
		masterModel, err = dao.GetUserCategory(db, nc.ID.Int64, 3, nc.Schema.String, data)
		if err.Error != nil {
			return
		}
		if data.UserCategory3ID == 0 {
			data.GroupID, err = dao.InsertNexsellerUserCategory(tx, nc.ID.Int64, mnID, 3, nc.Schema.String, &masterModel)
		} else {
			err = dao.UpdateNexsellerUserCategory(tx, data.UserCategory3ID, nc.Schema.String, &masterModel)
		}
		if err.Error != nil {
			return
		}
	}

	return
}
