package domain

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/PointOfSale/dto"
	"time"
)

type Brand struct {
	Id          string       `json:"id" db:"id" goqu:"skipinsert,skipupdate"`
	Name        string       `json:"name" db:"name"`
	Description string       `json:"description" db:"description"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at" goqu:"skipinsert,skipupdate"`
	UpdatedAt   sql.NullTime `json:"updated_at" db:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

type BrandRepository interface {
	Save(ctx context.Context, brand *Brand) error
	Update(ctx context.Context, brand *Brand, id string) error
	FindById(ctx context.Context, id string) (brand Brand, err error)
	FindAll(ctx context.Context) (brands []Brand, err error)
	Delete(ctx context.Context, id string) error
}

type BrandService interface {
	Save(ctx context.Context, req dto.CreateBrandRequest) error
	Update(ctx context.Context, req dto.UpdateBrandRequest, id string) error
	Index(ctx context.Context) ([]dto.BrandData, error)
	GetById(ctx context.Context, id string) (dto.BrandData, error)
	Delete(ctx context.Context, req string) error
}

type BrandHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Index(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
}
