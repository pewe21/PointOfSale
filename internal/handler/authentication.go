package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/PointOfSale/dto"
	"github.com/pewe21/PointOfSale/internal/domain"
	"github.com/pewe21/PointOfSale/internal/response"
	"net/http"
	"time"
)

type authenticationHandler struct {
	domain.AuthService
}

func NewAuthenticationHandler(authService domain.AuthService) domain.AuthHandler {
	return &authenticationHandler{AuthService: authService}
}

func (h authenticationHandler) SignIn(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 5*time.Second)
	defer cancel()
	var req dto.SignInRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ResponseError(err.Error(), http.StatusBadRequest))
	}
	token, err := h.AuthService.SignIn(c, req)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ResponseError(err.Error(), http.StatusUnauthorized))
	}

	resp := dto.SignInResponse{Token: token}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(resp))
}
