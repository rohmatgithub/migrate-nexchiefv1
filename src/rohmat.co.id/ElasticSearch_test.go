package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olivere/elastic"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ElasticClient *elastic.Client

func TestCheckExistIndex(t *testing.T) {
	setConnectionElastic(t)
	exist, err := elastic.NewIndicesExistsService(ElasticClient).Index([]string{"nc.param"}).Do(context.Background())
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	fmt.Println(exist)
}

func TestSettingMapping(t *testing.T) {
	index := "test_botol"
	mappings := `{
   "settings":{
      "analysis":{
         "analyzer":{
            "my_analyzer":{ 
               "type":"custom",
               "tokenizer":"whitespace",
               "filter":[
                  "lowercase"
               ]
            }
         }
      }
   },
   "mappings":{
       "properties":{
          "name": {
             "type":"text",
             "analyzer":"my_analyzer"
         },
         "email": {
             "type":"text",
             "analyzer":"my_analyzer"
         }
      }
   }
}`
	setConnectionElastic(t)
	ctx := context.Background()
	exist, err := elastic.NewIndicesExistsService(ElasticClient).Index([]string{index}).Do(ctx)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	if !exist {
		// create index
		var create *elastic.IndicesCreateResult
		create, err = ElasticClient.CreateIndex(index).Body(mappings).Do(ctx)
		if err != nil {
			assert.FailNow(t, err.Error())
		} else {
			fmt.Println("CreateIndex():", create)
		}
	}

}

func searchElastic(t *testing.T) {

}
func setConnectionElastic(t *testing.T) {
	var err error
	ElasticClient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	fmt.Println("elastic search connected")
}

func TestInsertMysql(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/debezium")

	// if there is an error opening the connection, handle it
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	//go func() {
	err = insertHeader(db)
	if err != nil {
		fmt.Println("err header : ", err)
	}
	//}()

	//go func() {
	err = insertItem(db)
	if err != nil {
		fmt.Println("err item : ", err)
	}
	//}()

	//time.Sleep(3 * time.Second)
}

func insertHeader(db *sql.DB) (err error) {
	query := "" +
		"INSERT INTO mustselllistsalesman " +
		"(mslSalesmanID, principalID, userLevel1ID, userLevel2ID, customerSubTypeID, salesmanTypeIDList, salesmanCategoryIDList) " +
		"VALUES " +
		"('MSLID7', 'ASW', 'GEO1', 'GEO2', 'GTRADE', 'TO`CV', 'TO`CV');"
	_, err = db.Exec(query)
	return
}

func insertItem(db *sql.DB) (err error) {
	query := " INSERT INTO mustselllistsalesmanitem " +
		"(mslSalesmanID, principalID, productCode, mslStatus, qtyMinimum) " +
		"VALUES " +
		"('MSLID7', 'ASW', 'CC-MGR-32', 'Mandatory', 10);"

	_, err = db.Exec(query)
	return
}

