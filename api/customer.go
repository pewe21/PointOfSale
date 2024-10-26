package api

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func NewCustomerApi(app *fiber.App, conn *sql.DB) {
	customer := InitializedCustomer(conn)
	group := app.Group("/customer")
	group.Get("/", customer.Index)
	group.Get("/:id", customer.GetById)
	group.Get("/username/:username", customer.GetByUsername)
	group.Get("/email/:email", customer.GetByEmail)
	group.Post("/", customer.Create)
	group.Post("/add_role", customer.AddRole)
	group.Put("/change_role/:customer_role", customer.ChangeRole)
	group.Put("/:id", customer.Update)
	group.Delete("/:id", customer.Delete)

}
