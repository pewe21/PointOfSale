package api

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func NewTypeApi(app *fiber.App, conn *sql.DB) {
	_type := InitializedType(conn)
	group := app.Group("/type")
	group.Get("/", _type.Index)
	group.Get("/:id", _type.GetById)
	group.Post("/", _type.Create)
	group.Put("/:id", _type.Update)
	group.Delete("/:id", _type.Delete)
}
