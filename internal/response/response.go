package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func ResSuccess(ctx *fiber.Ctx, successType int, data interface{}) error {
	switch successType {
	case http.StatusCreated:
		return ctx.Status(http.StatusCreated).JSON(ResponseCreateSuccess())
	case http.StatusOK:
		return ctx.Status(http.StatusOK).JSON(ResponseSuccess(data))
	default:
		return ctx.Status(http.StatusOK).JSON(ResponseSuccess(data))
	}
}
