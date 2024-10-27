package main

import (
	"database/sql"

	"github.com/pewe21/PointOfSale/internal/config"
	"github.com/pewe21/PointOfSale/internal/database"
)

func InitializedLoader() *config.Config {

	return &config.Config{
		Server: config.Server{
			Host: "127.0.0.1",
			Port: "3000",
		},
		Database: config.Database{
			Host:     "localhost",
			Port:     "5432",
			User:     "pos",
			Password: "passpos",
			Name:     "test_pointOfSale",
			Tz:       "Asia/Jakarta",
		},
		Jwt: config.Jwt{
			Secret: "secret",
			Exp:    60,
		},
	}
}

func GlobalSetupTest() *sql.DB {
	conf := InitializedLoader()

	conn := database.InitDB(conf.Database, false)
	return conn
}
