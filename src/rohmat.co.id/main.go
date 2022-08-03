package main

import (
	_ "github.com/golang-migrate/migrate/database/postgres"
	"rohmat.co.id/config"
	"rohmat.co.id/serverconfig"

	//_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)
func main() {
	config.GenerateConfiguration()
	serverconfig.SetServerAttribute()


}