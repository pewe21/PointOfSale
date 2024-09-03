package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/PointOfSale/internal/domain"
	"github.com/pewe21/PointOfSale/internal/response"
	"net/http"
	"time"
)

type handlerProduct struct {
	service domain.ProductService
}

func NewHandlerProduct(service domain.ProductService) domain.ProductHandler {
	return &handlerProduct{service: service}
}

func (h handlerProduct) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	products, err := h.service.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(products))
}

func (h handlerProduct) GetById(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	id := ctx.Params("id")
	product, err := h.service.GetById(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(product))
}

func (h handlerProduct) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	product := domain.Product{}
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	err := h.service.Create(c, &product)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))

	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseCreateSuccess())
}

func (h handlerProduct) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	id := ctx.Params("id")
	product := domain.Product{}
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	err := h.service.Update(c, &product, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))

	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(""))
}

func (h handlerProduct) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	id := ctx.Params("id")

	err := h.service.Delete(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))

	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(""))
}
