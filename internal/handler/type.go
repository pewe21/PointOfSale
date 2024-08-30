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

type typeHandler struct {
	service domain.TypeService
}

func NewHandlerType(service domain.TypeService) domain.TypeHandler {
	return &typeHandler{service: service}
}

func (h typeHandler) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	var req dto.CreateTypeRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	if err := h.service.Save(c, req); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusCreated).JSON(response.ResponseCreateSuccess())

}

func (h typeHandler) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	id := ctx.Params("id")
	var req dto.UpdateTypeRequest

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

func (h typeHandler) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	_type, err := h.service.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(_type))
}

func (h typeHandler) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	id := ctx.Params("id")

	err := h.service.Delete(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(""))
}

func (h typeHandler) GetById(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	id := ctx.Params("id")
	_type, err := h.service.GetById(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(_type))
}
