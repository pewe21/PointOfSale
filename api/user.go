package api

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func NewUserApi(app *fiber.App, conn *sql.DB) {
	user := InitializedUser(conn)

	group := app.Group("/user")
	group.Get("/", user.Index)
	group.Post("/", user.Create)
	group.Put("/:id", user.Update)
	group.Delete("/:id", user.Delete)

}
