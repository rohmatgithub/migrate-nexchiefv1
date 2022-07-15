package service

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"rohmat.co.id/dao"
	"rohmat.co.id/model"
	"testing"
	"time"
)

func TestInsertUserLevel1(t *testing.T) {
	path := "C:\\cdc-tools\\data sql\\local\\user-level-1.json"

	StartReadFile(path, saveUserLevel1)
	time.Sleep(5 * time.Second)
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

func TestInsertUserLeve2(t *testing.T) {
	path := "C:\\cdc-tools\\data sql\\local\\user-level-2.json"

	StartReadFile(path, saveUserLevel2)
	time.Sleep(5 * time.Second)
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

func TestInsertUserLeve3(t *testing.T) {
	path := "C:\\cdc-tools\\data sql\\local\\user-level-3.json"

	StartReadFile(path, saveUserLevel3)
	time.Sleep(5 * time.Second)
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

func TestInsertUserLeve4(t *testing.T) {
	path := "C:\\cdc-tools\\data sql\\local\\user-level-4.json"

	StartReadFile(path, saveUserLevel4)
	time.Sleep(5 * time.Second)

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

func TestInsertUserLeve5(t *testing.T) {
	path := "C:\\cdc-tools\\data sql\\local\\user-level-5.json"

	StartReadFile(path, saveUserLevel5)
	time.Sleep(5 * time.Second)

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
//func TestQueryParent(t *testing.T) {
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
