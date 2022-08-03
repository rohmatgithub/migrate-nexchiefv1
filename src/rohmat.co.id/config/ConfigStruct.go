package config

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"os"
)

var ApplicationConfiguration Configuration

type Configuration interface {
	GetPostgreSQLAddress() string
	GetPostgreSQLDefaultSchema() string
	GetPostgreSQLMaxOpenConnection() int
	GetPostgreSQLMaxIdleConnection() int
	GetDirPath() DirPath
}

func GenerateConfiguration() {
	var err error
	enviName := os.Getenv("NexchiefDbzConfig")
	temp := DevelopmentConfig{}
	err = gonfig.GetConf(enviName+"/config_development.json", &temp)
	ApplicationConfiguration = &temp

	if err != nil {
		fmt.Print(err)
		os.Exit(2)
	}
}
