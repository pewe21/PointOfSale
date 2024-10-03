package api

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func NewProductApi(app *fiber.App, conn *sql.DB) {
	product := InitializedProduct(conn)
	group := app.Group("/product")
	group.Get("/", product.Index)
	group.Get("/:id", product.GetById)
	group.Post("/", product.Create)
	group.Put("/:id", product.Update)
	group.Delete("/:id", product.Delete)
}
