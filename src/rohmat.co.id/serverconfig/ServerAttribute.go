package serverconfig

import (
	"bufio"
	"database/sql"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/bukalapak/go-redis"
	"nexsoft.co.id/nexcommon/db/dbconfig"
	"os"
	"rohmat.co.id/config"
	"strconv"
)

var ServerAttribute serverAttribute

type serverAttribute struct {
	Version                                  string
	DBConnection                             *sql.DB
	SQLMigrationResolutionDir                string
	Write *bufio.Writer
}

func SetServerAttribute() {
	dbParam := config.ApplicationConfiguration.GetPostgreSQLDefaultSchema()
	dbConnection := config.ApplicationConfiguration.GetPostgreSQLAddress()
	dbMaxOpenConnection := config.ApplicationConfiguration.GetPostgreSQLMaxOpenConnection()
	dbMaxIdleConnection := config.ApplicationConfiguration.GetPostgreSQLMaxIdleConnection()
	ServerAttribute.DBConnection = dbconfig.GetDbConnection(dbParam, dbConnection, dbMaxOpenConnection, dbMaxIdleConnection)
	f, _ := os.Create("file.log")
	ServerAttribute.Write = bufio.NewWriter(f)
}

func getRedisClient(host string, port int, db int, password string, optCB *hystrix.CommandConfig) *redis.Client {
	redisAddress := host + ":" + strconv.Itoa(port)
	opts := &redis.Options{
		CircuitBreaker: optCB,
		Addr:           redisAddress,
		Password:       password,
		DB:             db,
	}

	return redis.NewClient(opts)
}
