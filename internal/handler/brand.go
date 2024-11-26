package handler

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/PointOfSale/dto"
	"github.com/pewe21/PointOfSale/internal/domain"
	"github.com/pewe21/PointOfSale/internal/response"
	"github.com/pewe21/PointOfSale/internal/util"
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
		return ctx.Status(http.StatusUnprocessableEntity).JSON(response.ResponseError(err.Error(), http.StatusUnprocessableEntity))
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		errs := err.(validator.ValidationErrors)
		return ctx.Status(http.StatusBadRequest).JSON(
			response.ResponseError(
				errs.Error(),
				http.StatusBadRequest,
			),
		)
	}

	if err := h.service.Save(c, req); err != nil {
		if strings.Contains(err.Error(), "already exist") {
			return ctx.Status(http.StatusConflict).JSON(response.ResponseError(err.Error(), http.StatusConflict))
		}
		// return response.ResError(ctx, http.StatusInternalServerError, nil)
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusCreated).JSON(response.ResponseCreateSuccess())
	// return response.ResSuccess(ctx, http.StatusCreated, nil)

}

func (h brandHandler) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	id := ctx.Params("id")

	var req dto.UpdateBrandRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(response.ResponseError(err.Error(), http.StatusUnprocessableEntity))
		// return response.ResError(ctx, http.StatusUnprocessableEntity, nil)
	}

	if errValid := util.Validate(req); errValid != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ResponseError(errValid.Error(), http.StatusBadRequest))
	}

	_, err := h.service.GetById(c, id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ResponseError(err.Error(), http.StatusBadRequest))
		// return response.ResError(ctx, http.StatusBadRequest, nil)
	}

	err = h.service.Update(c, req, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
		// return response.ResError(ctx, http.StatusInternalServerError, nil)
	}

	// return response.ResSuccess(ctx, http.StatusOK, nil)
	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(""))
}

func (h brandHandler) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	brands, err := h.service.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
		// return response.ResError(ctx, http.StatusInternalServerError, nil)
	}

	return response.ResSuccess(ctx, http.StatusOK, brands)
}

func (h brandHandler) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	id := ctx.Params("id")

	err := h.service.Delete(c, id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return ctx.Status(http.StatusNotFound).JSON(response.ResponseError(err.Error(), http.StatusNotFound))
			// return response.ResError(ctx, http.StatusNotFound, nil)
		}
		// return response.ResError(ctx, http.StatusInternalServerError, nil)
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	// return response.ResSuccess(ctx, http.StatusOK, "")
	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(""))
}

func (h brandHandler) GetById(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	id := ctx.Params("id")
	brand, err := h.service.GetById(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(brand))
}
