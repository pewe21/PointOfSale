package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func InitializedLoader() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtExp, err := strconv.Atoi(os.Getenv("JWT_EXP"))
	if err != nil {
		log.Fatal(err)
	}

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: Database{
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Name:     os.Getenv("DATABASE_NAME"),
			Tz:       os.Getenv("DATABASE_TZ"),
		},
		Jwt: Jwt{
			Secret: os.Getenv("JWT_SECRET"),
			Exp:    jwtExp,
		},
		Redis: Redis{
			Host: os.Getenv("REDIS_HOST"),
			Port: os.Getenv("REDIS_PORT"),
		},
	}
}
