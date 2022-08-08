package main

import (
	"fmt"
	_ "github.com/golang-migrate/migrate/database/postgres"
	"rohmat.co.id/config"
	"rohmat.co.id/serverconfig"
	"rohmat.co.id/service"
	"sync"
	//_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

func main() {
	config.GenerateConfiguration()
	serverconfig.SetServerAttribute()

	//// geo tree
	//service.StartSaveUserLevel1()
	//service.StartSaveUserLevel2()
	//service.StartSaveUserLevel3()
	//service.StartSaveUserLevel4()
	//service.StartSaveUserLevel5()
	//
	//// save division
	//service.StartSaveDivision()
	//
	//// save mapping distributor
	//service.StartSaveMappingDistributor()
	//service.StartSaveDistributor()
	//
	//// save vendor
	//service.StartSaveVendor()
	//
	service.StartSaveIsland()

	// save salesman
	service.StartInsertSalesman()


	var wg sync.WaitGroup

	service.StartSaveNexsellerProduct(&wg)
	// save nexseller customer
	service.StartSaveNexsellerCustomer(&wg)

	wg.Wait()
	fmt.Println("====== FINISH ======")
	//time.Sleep(10 * time.Second)
}
