package api

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func NewBrandApi(app *fiber.App, conn *sql.DB) {
	brand := InitializedBrand(conn)
	group := app.Group("/brand")
	group.Get("/", brand.Index)
	group.Get("/:id", brand.GetById)
	group.Post("/", brand.Create)
	group.Put("/:id", brand.Update)
	group.Delete("/:id", brand.Delete)
}
