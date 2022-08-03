package serverconfig

import (
	"database/sql"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/bukalapak/go-redis"
	"nexsoft.co.id/nexcommon/db/dbconfig"
	"rohmat.co.id/config"
	"strconv"
)

var ServerAttribute serverAttribute

type serverAttribute struct {
	Version                                  string
	DBConnection                             *sql.DB
	SQLMigrationResolutionDir                string
}

func SetServerAttribute() {
	dbParam := config.ApplicationConfiguration.GetPostgreSQLDefaultSchema()
	dbConnection := config.ApplicationConfiguration.GetPostgreSQLAddress()
	dbMaxOpenConnection := config.ApplicationConfiguration.GetPostgreSQLMaxOpenConnection()
	dbMaxIdleConnection := config.ApplicationConfiguration.GetPostgreSQLMaxIdleConnection()
	ServerAttribute.DBConnection = dbconfig.GetDbConnection(dbParam, dbConnection, dbMaxOpenConnection, dbMaxIdleConnection)
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
