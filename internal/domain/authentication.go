package domain

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/PointOfSale/dto"
)

type AuthRepository interface {
	Find(ctx context.Context, id string) (user User, err error)
}

type AuthService interface {
	SignIn(ctx context.Context, req dto.SignInRequest) (string, error)
}

type AuthHandler interface {
	SignIn(ctx *fiber.Ctx) error
}
