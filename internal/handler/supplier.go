package handler

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/PointOfSale/dto"
	"github.com/pewe21/PointOfSale/internal/domain"
	"github.com/pewe21/PointOfSale/internal/response"
	"github.com/pewe21/PointOfSale/internal/util"
)

type supplierHandler struct {
	service domain.SupplierService
}

func NewHandlerSupplier(service domain.SupplierService) domain.SupplierHandler {
	return &supplierHandler{service: service}
}

func (h supplierHandler) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	var req dto.CreateSupplierRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
		// return response.ResError(ctx, http.StatusUnprocessableEntity, nil)
	}

	if errValid := util.Validate(req); errValid != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ResponseError(errValid.Error(), http.StatusBadRequest))
	}

	if err := h.service.Save(c, req); err != nil {
		if strings.Contains(err.Error(), "already exist") {
			return ctx.Status(http.StatusConflict).JSON(response.ResponseError(err.Error(), http.StatusConflict))
			// return response.ResError(ctx, http.StatusConflict, err.Error())
		}
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
		// return response.ResError(ctx, http.StatusInternalServerError, nil)
	}
	return ctx.Status(http.StatusCreated).JSON(response.ResponseCreateSuccess())
}

func (h supplierHandler) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	id := ctx.Params("id")
	var req dto.UpdateSupplierRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	if errValid := util.Validate(req); errValid != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ResponseError(errValid.Error(), http.StatusBadRequest))
	}

	_, err = h.service.GetById(c, id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ResponseError(err.Error(), http.StatusBadRequest))
	}

	err = h.service.Update(c, req, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(""))
}

func (h supplierHandler) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	customers, err := h.service.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(customers))
}

func (h supplierHandler) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	id := ctx.Params("id")

	err := h.service.Delete(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(""))
}

func (h supplierHandler) GetById(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	id := ctx.Params("id")
	customer, err := h.service.GetById(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(customer))
}
