package api

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/PointOfSale/internal/config"
)

func NewAuthApi(app *fiber.App, conn *sql.DB, cnf *config.Jwt) {
	authentication := InitializedAuthentication(conn, cnf)
	group := app.Group("/auth")
	group.Post("/", authentication.SignIn)
}
