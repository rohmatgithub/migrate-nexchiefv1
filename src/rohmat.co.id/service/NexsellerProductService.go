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
	"sync"
)


func StartSaveNexsellerProduct(wg *sync.WaitGroup) {
	// 2021-01-01
	path := config.ApplicationConfiguration.GetDirPath().PathDir +
		config.ApplicationConfiguration.GetDirPath().NexsellerProduct

	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, file := range files {
		if !file.IsDir() {
			wg.Add(1)
			go doStartWaitGroupNexProduct(wg, path+"/"+file.Name(), fmt.Sprintf("nexseller product %d", i+1))
		}
	}
}

func doStartWaitGroupNexProduct(wg *sync.WaitGroup, path string, partData string) {
	defer wg.Done()

	StartReadFile(path, SaveNexsellerProduct, partData)

}

func SaveNexsellerProduct(db *sql.DB, dataByte []byte)(errorModel model.ErrorModel) {
	var (
		nexchiefAccount model.NexchiefAccount
		data            model.NexsellerProduct
		mnID            int64
	)
	_ = json.Unmarshal(dataByte, &data)

	nexchiefAccount, errorModel = dao.GetNexchiefAccountID(db, data.NcCode)
	if errorModel.Error != nil {
		return
	}
	// get mapping nexseller
	mnID, errorModel = dao.GetMappingNexsellerID(db, nexchiefAccount.ID.Int64, data.MnCode)
	if errorModel.Error != nil {
		return
	}
	errorModel = dao.GetNexProductFK(db, nexchiefAccount.Schema.String, mnID, &data)
	if errorModel.Error != nil {
		return
	}

	pkString := data.NcCode + data.MnCode + data.Code
	data.PkChecksum = util.CheckSumWithXXHASH([]byte(pkString))
	if data.ID == 0 {
		// insert
		errorModel = dao.InsertNexsellerProduct(db, nexchiefAccount, mnID, &data)
	} else {
		errorModel = dao.UpdateNexsellerProduct(db, nexchiefAccount, mnID, &data)
	}
	return
}
