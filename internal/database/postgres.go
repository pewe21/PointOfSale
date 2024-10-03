package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pewe21/PointOfSale/internal/config"
	"log"
	"time"
)

func InitDB(conf config.Database, setLimits bool) *sql.DB {

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
	log.Printf("PINGING SUCCESS")
	return db
}
