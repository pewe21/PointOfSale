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

type brandHandler struct {
	service domain.BrandService
}

func NewHandlerType(service domain.BrandService) domain.BrandHandler {
	return &brandHandler{service: service}
}

func (h brandHandler) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	var req dto.CreateBrandRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	if err := h.service.Save(c, req); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusCreated).JSON(response.ResponseCreateSuccess())

}

func (h brandHandler) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	id := ctx.Params("id")

	var req dto.UpdateBrandRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	_, err := h.service.GetById(c, id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ResponseError(err.Error(), http.StatusBadRequest))
	}

	err = h.service.Update(c, req, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(""))
}

func (h brandHandler) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	_type, err := h.service.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(_type))
}

func (h brandHandler) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	id := ctx.Params("id")

	err := h.service.Delete(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(""))
}

func (h brandHandler) GetById(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	id := ctx.Params("id")
	_type, err := h.service.GetById(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(_type))
}
