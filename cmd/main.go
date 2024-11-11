package main

import (
	"log"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/PointOfSale/api"
	"github.com/pewe21/PointOfSale/internal/config"
	"github.com/pewe21/PointOfSale/internal/database"
)

func main() {
	conf := config.InitializedLoader()

	conn := database.InitDB(conf.Database, false)

	//redis := cache.NewRedisCache(conf.Redis)

	app := fiber.New()
	//auth
	api.NewAuthApi(app, conn, &conf.Jwt)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(conf.Jwt.Secret)},
	}))
	api.NewUserApi(app, conn)
	api.NewCustomerApi(app, conn)
	api.NewSupplierApi(app, conn)
	api.NewBrandApi(app, conn)
	api.NewProductApi(app, conn)
	api.NewRoleApi(app, conn)

	err := app.Listen(conf.Server.Host + ":" + conf.Server.Port)
	if err != nil {
		log.Fatal("Application failed to start")
	}
}
