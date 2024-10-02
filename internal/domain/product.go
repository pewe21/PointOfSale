package domain

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/PointOfSale/dto"
	"time"
)

type Product struct {
	Id         string       `json:"id" db:"id" goqu:"skipinsert,skipupdate"`
	Name       string       `json:"name" db:"name"`
	SKU        string       `json:"sku" db:"sku"`
	Stock      int          `json:"stock" db:"stock"`
	BrandId    string       `json:"brand_id" db:"brand_id"`
	SupplierId string       `json:"supplier_id" db:"supplier_id"`
	CreatedAt  time.Time    `json:"created_at" db:"created_at" goqu:"skipinsert,skipupdate"`
	UpdatedAt  sql.NullTime `json:"updated_at" db:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

type ProductWithDetail struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	SKU          string `json:"sku"`
	Stock        int    `json:"stock"`
	BrandId      string `json:"brand_id" db:"brand_id"`
	BrandName    string `json:"brand_name" db:"brand_name"`
	SupplierId   string `json:"supplier_id" db:"supplier_id"`
	SupplierName string `json:"supplier_name" db:"supplier_name"`
}

type ProductRepository interface {
	FindAll(ctx context.Context) (products []ProductWithDetail, err error)
	FindById(ctx context.Context, id string) (product ProductWithDetail, err error)
	Save(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product, id string) error
	Delete(ctx context.Context, id string) error
}
type ProductService interface {
	Index(ctx context.Context) (products []ProductWithDetail, err error)
	IndexNew(ctx context.Context) (productsx []dto.ProductxDto, err error)
	GetById(ctx context.Context, id string) (product dto.ProductxDto, err error)
	Create(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product, id string) error
	Delete(ctx context.Context, id string) error
}

type ProductHandler interface {
	Index(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
