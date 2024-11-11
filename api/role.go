package api

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func NewRoleApi(app *fiber.App, conn *sql.DB) {
	role := InitializedRole(conn)
	group := app.Group("role")
	group.Get("/", role.Index)
	group.Get("/:id", role.GetById)
	group.Post("/", role.Create)
	group.Put("/:id", role.Update)
	group.Delete("/:id", role.Delete)

}
