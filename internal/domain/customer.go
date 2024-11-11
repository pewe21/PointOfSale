package domain

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/PointOfSale/dto"
	"time"
)

type Customer struct {
	Id        string       `json:"id" db:"id" goqu:"skipinsert,skipupdate"`
	Name      string       `json:"name" db:"name"`
	Username  string       `json:"username" db:"username"`
	Password  string       `json:"password" db:"password"`
	Email     string       `json:"email" db:"email"`
	Phone     string       `json:"phone" db:"phone"`
	Address   string       `json:"address" db:"address"`
	Verified  *time.Time   `json:"verified" db:"verified"`
	CreatedAt time.Time    `json:"created_at" db:"created_at" goqu:"skipinsert"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at" `
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at" `
}

type CustomerRepository interface {
	Save(ctx context.Context, customer *Customer) error
	Update(ctx context.Context, customer *Customer, id string) error
	FindById(ctx context.Context, id string) (customer Customer, err error)
	FindByUsername(ctx context.Context, username string) (customer Customer, err error)
	FindByEmail(ctx context.Context, email string) (customer Customer, err error)
	FindAll(ctx context.Context) (customers []Customer, err error)
	Delete(ctx context.Context, id string) error
}

type CustomerService interface {
	Save(ctx context.Context, req dto.CreateCustomerRequest) error
	Update(ctx context.Context, req dto.UpdateCustomerRequest, id string) error
	Index(ctx context.Context) ([]dto.CustomerData, error)
	GetById(ctx context.Context, id string) (customer Customer, err error)
	GetByUsername(ctx context.Context, username string) (customer Customer, err error)
	GetByEmail(ctx context.Context, email string) (customer Customer, err error)
	Delete(ctx context.Context, req string) error
	AddRole(ctx context.Context, req dto.AddCustomerRoleRequest) error
	ChangeRole(ctx context.Context, req dto.UpdateCustomerRoleRequest, customerId string) error
}

type CustomerHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	GetByUsername(ctx *fiber.Ctx) error
	GetByEmail(ctx *fiber.Ctx) error
	Index(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	AddRole(ctx *fiber.Ctx) error
	ChangeRole(ctx *fiber.Ctx) error
}
