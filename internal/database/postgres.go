package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pewe21/PointOfSale/internal/config"
	"log"
)

func InitDB(conf config.Database) *sql.DB {
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

	err = db.Ping()

	if err != nil {
		log.Fatal("failed to ping database")
	}
	log.Printf("PINGING SUCCESS")
	return db
}
