package main

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/PointOfSale/api"
	"github.com/pewe21/PointOfSale/internal/config"
	"github.com/pewe21/PointOfSale/internal/database"
	"log"
)

func main() {
	conf := config.InitializedLoader()

	conn := database.InitDB(conf.Database)

	//redis := cache.NewRedisCache(conf.Redis)

	app := fiber.New()
	//auth
	api.NewAuthApi(app, conn, &conf.Jwt)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(conf.Jwt.Secret)},
	}))
	api.NewUserApi(app, conn)

	err := app.Listen(conf.Server.Host + ":" + conf.Server.Port)
	if err != nil {
		log.Fatal("Application failed to start")
	}
}
