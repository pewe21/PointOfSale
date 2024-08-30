package api

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func NewSupplierApi(app *fiber.App, conn *sql.DB) {
	supplier := InitializedSupplier(conn)
	group := app.Group("/supplier")
	group.Get("/", supplier.Index)
	group.Get("/:id", supplier.GetById)
	group.Post("/", supplier.Create)
	group.Put("/:id", supplier.Update)
	group.Delete("/:id", supplier.Delete)
}
