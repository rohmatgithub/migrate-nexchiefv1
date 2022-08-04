package main

import (
	_ "github.com/golang-migrate/migrate/database/postgres"
	"rohmat.co.id/config"
	"rohmat.co.id/serverconfig"
	"rohmat.co.id/service"
	"time"

	//_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

func main() {
	config.GenerateConfiguration()
	serverconfig.SetServerAttribute()

	// geo tree
	service.StartSaveUserLevel1()
	service.StartSaveUserLevel2()
	service.StartSaveUserLevel3()
	service.StartSaveUserLevel4()
	service.StartSaveUserLevel5()

	// save division
	service.StartSaveDivision()

	// save mapping distributor
	service.StartSaveMappingDistributor()
	service.StartSaveDistributor()

	// save vendor
	service.StartSaveVendor()

	// save salesman
	service.StartInsertSalesman()

	// save nexseller customer
	service.StartSaveNexsellerCustomer()
	time.Sleep(10 * time.Second)
}
