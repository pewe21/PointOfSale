package domain

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/PointOfSale/dto"
	"time"
)

type Type struct {
	Id          string       `json:"id" db:"id" goqu:"skipinsert"`
	Name        string       `json:"name" db:"name"`
	Description string       `json:"description" db:"description"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at" goqu:"skipinsert"`
	UpdatedAt   sql.NullTime `json:"updated_at" db:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

type TypeRepository interface {
	Save(ctx context.Context, _type *Type) error
	Update(ctx context.Context, _type *Type, id string) error
	FindById(ctx context.Context, id string) (_type Type, err error)
	FindAll(ctx context.Context) (types []Type, err error)
	Delete(ctx context.Context, id string) error
}

type TypeService interface {
	Save(ctx context.Context, req dto.CreateTypeRequest) error
	Update(ctx context.Context, req dto.UpdateTypeRequest, id string) error
	Index(ctx context.Context) ([]dto.TypeData, error)
	GetById(ctx context.Context, id string) (_type dto.TypeData, err error)
	Delete(ctx context.Context, req string) error
}

type TypeHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Index(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
}
