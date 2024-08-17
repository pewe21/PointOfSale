package domain

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/PointOfSale/dto"
	"time"
)

type User struct {
	Id        string       `json:"id" db:"id" goqu:"skipinsert"`
	Name      string       `json:"name" db:"name"`
	Email     string       `json:"email" db:"email"`
	Password  string       `json:"password" db:"password"`
	Phone     string       `json:"phone" db:"phone"`
	CreatedAt time.Time    `json:"created_at" db:"created_at" goqu:"skipinsert"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User, id string) error
	FindAll(ctx context.Context) (result []User, err error)
	FindById(ctx context.Context, id string) (user User, err error)
	Delete(ctx context.Context, id string) error
	FindByEmail(ctx context.Context, email string) (user User, err error)
}

type UserService interface {
	Index(ctx context.Context) ([]dto.UserData, error)
	Save(ctx context.Context, req dto.CreateUserRequest) error
	Update(ctx context.Context, req dto.UpdateUserRequest, id string) error
	Delete(ctx context.Context, req string) error
}

type UserHandler interface {
	Index(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
