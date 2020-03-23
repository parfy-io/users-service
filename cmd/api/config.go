package main

import (
	"fmt"
	"github.com/ardanlabs/conf"
	"os"
)

type config struct {
	ServerAddress string `conf:"help:Server-address network-interface to bind on e.g.: '127.0.0.1:8080',default:0.0.0.0:80"`
	DB            struct {
		Host                 string `conf:"help:Database-Host,required"`
		Port                 int    `conf:"help:Database-Port,default:5432"`
		Name                 string `conf:"help:Database-name,default:'users-service'"`
		Username             string `conf:"help:Database-Username"`
		Password             string `conf:"help:Database-Password,noprint"`
		MigrationsFolderPath string `conf:"help:Database Migrations Folder Path,default:/db-migrations"`
	}
}

func newConfig() (config, error) {
	cfg := config{}

	if origErr := conf.Parse(os.Environ(), "US", &cfg); origErr != nil {
		usage, err := conf.Usage("US", &cfg)
		if err != nil {
			return cfg, err
		}
		fmt.Println(usage)
		return cfg, origErr
	}

	return cfg, nil
}
