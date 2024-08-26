package domain

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/PointOfSale/dto"
	"time"
)

type Supplier struct {
	Id        string       `json:"id" db:"id" goqu:"skipinsert"`
	Name      string       `json:"name" db:"name"`
	Email     string       `json:"email" db:"email"`
	Address   string       `json:"address" db:"address"`
	Phone     string       `json:"phone" db:"phone"`
	CreatedAt time.Time    `json:"created_at" db:"created_at" goqu:"skipinsert"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

type SupplierRepository interface {
	Save(ctx context.Context, customer *Supplier) error
	Update(ctx context.Context, customer *Supplier, id string) error
	FindById(ctx context.Context, id string) (supplier Supplier, err error)
	FindAll(ctx context.Context) (suppliers []Supplier, err error)
	Delete(ctx context.Context, id string) error
}

type SupplierService interface {
	Save(ctx context.Context, req dto.CreateSupplierRequest) error
	Update(ctx context.Context, req dto.UpdateSupplierRequest, id string) error
	Index(ctx context.Context) ([]dto.SupplierData, error)
	GetById(ctx context.Context, id string) (supplier dto.SupplierData, err error)
	Delete(ctx context.Context, req string) error
}

type SupplierHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Index(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
}
