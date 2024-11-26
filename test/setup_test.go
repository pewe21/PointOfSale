package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"

	"github.com/pewe21/PointOfSale/internal/config"
)

type FetchedResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

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

func CreateAdminTest(conn *sql.DB) {

	pw, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	dataInsert := fmt.Sprintf("insert into users (email,name,password,phone) values ('admin@admin.com','test','%s','081123123123')", string(pw))

	_, err = conn.Exec(dataInsert)
	if err != nil {
		log.Fatal(err)
	}

}
