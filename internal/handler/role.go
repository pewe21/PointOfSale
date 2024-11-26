package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/PointOfSale/dto"
	"github.com/pewe21/PointOfSale/internal/domain"
	"github.com/pewe21/PointOfSale/internal/response"
	"github.com/pewe21/PointOfSale/internal/util"
)

type handlerRole struct {
	service domain.RoleService
}

func NewHandlerRole(service domain.RoleService) domain.RoleHandler {
	return &handlerRole{service: service}
}

func (h handlerRole) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	var req dto.CreateRoleRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	if errValid := util.Validate(req); errValid != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ResponseError(errValid.Error(), http.StatusBadRequest))
	}

	err := h.service.Save(c, req)

	if err != nil {
		if err.Error() == "role already exist" {
			return ctx.Status(http.StatusConflict).JSON(response.ResponseError(err.Error(), http.StatusConflict))
		}
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusCreated).JSON(response.ResponseCreateSuccess())
}

func (h handlerRole) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	var req dto.UpdateRoleRequest

	id := ctx.Params("id")

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	if errValid := util.Validate(req); errValid != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ResponseError(errValid.Error(), http.StatusBadRequest))
	}

	if err := h.service.Update(c, req, id); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			response.ResponseError(
				err.Error(),
				http.StatusInternalServerError,
			),
		)
	}

	return ctx.Status(http.StatusOK).JSON(
		response.ResponseSuccess(""),
	)
}

func (h handlerRole) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	datas, err := h.service.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			response.ResponseError(
				err.Error(),
				http.StatusInternalServerError,
			),
		)
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(datas))
}

func (h handlerRole) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	id := ctx.Params("id")

	if err := h.service.Delete(c, id); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			response.ResponseError(
				err.Error(),
				http.StatusInternalServerError,
			),
		)
	}

	return ctx.Status(http.StatusOK).JSON(
		response.ResponseSuccess("delete role successfully"),
	)
}

func (h handlerRole) GetById(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	id := ctx.Params("id")

	data, err := h.service.GetById(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			response.ResponseError(
				err.Error(),
				http.StatusInternalServerError,
			),
		)
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(data))
}
