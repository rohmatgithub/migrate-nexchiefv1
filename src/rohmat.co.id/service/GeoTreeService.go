package service

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"rohmat.co.id/config"
	"rohmat.co.id/dao"
	"rohmat.co.id/model"
)

func StartSaveUserLevel1() {
	path := config.ApplicationConfiguration.GetDirPath().PathDir +
		config.ApplicationConfiguration.GetDirPath().UserLevel1

	StartReadFile(path, saveUserLevel1, "user level 1")
}

func saveUserLevel1(db *sql.DB, dataByte []byte) (err model.ErrorModel) {
	var data model.DataUserLevel
	_ = json.Unmarshal(dataByte, &data)

	nexchiefAccount, err := dao.GetNexchiefAccountID(db, data.NcCode)
	if err.Error != nil {
		return
	}
	data.Level = 1
	data.Code = data.Code1
	data.ParentID = 0

	err = dao.GetGeoTreeID(db, nexchiefAccount.ID.Int64, &data)
	if err.Error != nil {
		return
	} else if data.ID > 0 {
		// update
		err = dao.UpdateGeoTree(db, nexchiefAccount.ID.Int64, &data)
	} else {
		err = dao.InsertGeoTree(db, nexchiefAccount.ID.Int64, &data)
	}
	if err.Error != nil {
		return
	}
	return
}

func StartSaveUserLevel2() {
	path := config.ApplicationConfiguration.GetDirPath().PathDir +
		config.ApplicationConfiguration.GetDirPath().UserLevel2

	StartReadFile(path, saveUserLevel2, "user level 2")
}

func saveUserLevel2(db *sql.DB, dataByte []byte) (err model.ErrorModel) {
	var data model.DataUserLevel
	_ = json.Unmarshal(dataByte, &data)

	nexchiefAccount, err := dao.GetNexchiefAccountID(db, data.NcCode)
	if err.Error != nil {
		return
	}
	data.Level = 2
	data.Code = data.Code2
	err = dao.GetParent(db, nexchiefAccount.ID.Int64, &data, []interface{}{data.Code1})
	if err.Error != nil {
		return
	}
	err = dao.GetGeoTreeID(db, nexchiefAccount.ID.Int64, &data)
	if err.Error != nil {
		return
	} else if data.ID > 0 {
		// update
		err = dao.UpdateGeoTree(db, nexchiefAccount.ID.Int64, &data)
	} else {
		err = dao.InsertGeoTree(db, nexchiefAccount.ID.Int64, &data)
	}
	if err.Error != nil {
		return
	}
	return
}

func StartSaveUserLevel3() {
	path := config.ApplicationConfiguration.GetDirPath().PathDir +
		config.ApplicationConfiguration.GetDirPath().UserLevel3

	StartReadFile(path, saveUserLevel3, "user level 3")
}

func saveUserLevel3(db *sql.DB, dataByte []byte) (err model.ErrorModel) {
	var data model.DataUserLevel
	_ = json.Unmarshal(dataByte, &data)

	nexchiefAccount, err := dao.GetNexchiefAccountID(db, data.NcCode)
	if err.Error != nil {
		return
	}
	data.Level = 3
	data.Code = data.Code3
	err = dao.GetParent(db, nexchiefAccount.ID.Int64, &data, []interface{}{data.Code1, data.Code2})
	if err.Error != nil {
		return
	}
	err = dao.GetGeoTreeID(db, nexchiefAccount.ID.Int64, &data)
	if err.Error != nil {
		return
	} else if data.ID > 0 {
		// update
		err = dao.UpdateGeoTree(db, nexchiefAccount.ID.Int64, &data)
	} else {
		err = dao.InsertGeoTree(db, nexchiefAccount.ID.Int64, &data)
	}
	if err.Error != nil {
		return
	}
	return
}

func StartSaveUserLevel4() {
	path := config.ApplicationConfiguration.GetDirPath().PathDir +
		config.ApplicationConfiguration.GetDirPath().UserLevel4

	StartReadFile(path, saveUserLevel4, "user level 4")

}

func saveUserLevel4(db *sql.DB, dataByte []byte) (err model.ErrorModel) {
	var data model.DataUserLevel
	_ = json.Unmarshal(dataByte, &data)

	nexchiefAccount, err := dao.GetNexchiefAccountID(db, data.NcCode)
	if err.Error != nil {
		return
	}
	data.Level = 4
	data.Code = data.Code4
	err = dao.GetParent(db, nexchiefAccount.ID.Int64, &data, []interface{}{data.Code1, data.Code2, data.Code3})
	if err.Error != nil {
		return
	}
	err = dao.GetGeoTreeID(db, nexchiefAccount.ID.Int64, &data)
	if err.Error != nil {
		return
	} else if data.ID > 0 {
		// update
		err = dao.UpdateGeoTree(db, nexchiefAccount.ID.Int64, &data)
	} else {
		err = dao.InsertGeoTree(db, nexchiefAccount.ID.Int64, &data)
	}
	if err.Error != nil {
		return
	}
	return
}

func StartSaveUserLevel5() {
	path := config.ApplicationConfiguration.GetDirPath().PathDir +
		config.ApplicationConfiguration.GetDirPath().UserLevel5

	StartReadFile(path, saveUserLevel5, "user level 5")

}

func saveUserLevel5(db *sql.DB, dataByte []byte) (err model.ErrorModel) {
	var data model.DataUserLevel
	_ = json.Unmarshal(dataByte, &data)

	nexchiefAccount, err := dao.GetNexchiefAccountID(db, data.NcCode)
	if err.Error != nil {
		return
	}
	data.Level = 5
	data.Code = data.Code5
	err = dao.GetParent(db, nexchiefAccount.ID.Int64, &data, []interface{}{data.Code1, data.Code2, data.Code3, data.Code4})
	if err.Error != nil {
		return
	}
	err = dao.GetGeoTreeID(db, nexchiefAccount.ID.Int64, &data)
	data.Node = true
	if err.Error != nil {
		return
	} else if data.ID > 0 {
		// update
		err = dao.UpdateGeoTree(db, nexchiefAccount.ID.Int64, &data)
	} else {
		err = dao.InsertGeoTree(db, nexchiefAccount.ID.Int64, &data)
	}
	if err.Error != nil {
		return
	}
	return
}

//func TestQueryParent() {
//	m := make(map[string]interface{})
//	m["mapping_nexseller_id"] = 1
//	query, param := GetQueryParent("geo_tree", "geo_tree", "code", m, []interface{}{"GEO1", "GEO2", "GEO3", ""}, 0)
//	fmt.Println(query)
//	fmt.Println(len(param))
//	fmt.Println(param)
//}

func funcTest(arr *[]string) {
	temp := *arr
	for i := 0; i < len(temp); i++ {
		temp[i] = temp[i] + "--"
	}
}
