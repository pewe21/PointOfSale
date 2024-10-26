package domain

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/PointOfSale/dto"
	"time"
)

type Role struct {
	Id          string       `json:"id" db:"id" goqu:"skipinsert,skipupdate"`
	Name        string       `json:"name" db:"name"`
	DisplayName string       `json:"display_name" db:"display_name"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at" goqu:"skipupdate"`
	UpdatedAt   sql.NullTime `json:"updated_at" db:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

type RoleRepository interface {
	Save(ctx context.Context, role *Role) error
	Update(ctx context.Context, role *Role, id string) error
	FindById(ctx context.Context, id string) (role Role, err error)
	FindAll(ctx context.Context) (roles []Role, err error)
	Delete(ctx context.Context, id string) error
}

type RoleService interface {
	Save(ctx context.Context, req dto.CreateRoleRequest) error
	Update(ctx context.Context, req dto.UpdateRoleRequest, id string) error
	Index(ctx context.Context) ([]dto.RoleData, error)
	GetById(ctx context.Context, id string) (dto.RoleData, error)
	Delete(ctx context.Context, req string) error
}

type RoleHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Index(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
}
