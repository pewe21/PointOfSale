package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"

	"github.com/pewe21/PointOfSale/internal/config"
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
	}
}

func InitDBTest(conf config.Database, setLimits bool) *sql.DB {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Password,
		conf.Name,
		conf.Tz,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect database :" + err.Error())
	}

	if setLimits {
		fmt.Println("setting limits")
		db.SetMaxOpenConns(5)
		db.SetMaxIdleConns(5)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)

	if err != nil {
		log.Fatal("failed to ping database")
	}
	return db
}

func GlobalSetupTest() *sql.DB {
	conf := InitializedLoader()

	conn := InitDBTest(conf.Database, false)
	return conn
}
