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

type handlerCustomer struct {
	domain.CustomerService
}

func NewHandlerCustomer(customerService domain.CustomerService) domain.CustomerHandler {
	return &handlerCustomer{CustomerService: customerService}
}

func (h handlerCustomer) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	var req dto.CreateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	if err := h.CustomerService.Save(c, req); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusCreated).JSON(response.ResponseCreateSuccess())
}

func (h handlerCustomer) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	id := ctx.Params("id")
	var req dto.UpdateCustomerRequest

	_, err := h.CustomerService.GetById(c, id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ResponseError(err.Error(), http.StatusBadRequest))
	}

	err = h.CustomerService.Update(c, req, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(""))
}

func (h handlerCustomer) GetById(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	id := ctx.Params("id")
	customer, err := h.CustomerService.GetById(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(customer))
}

func (h handlerCustomer) GetByUsername(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	username := ctx.Params("username")
	customer, err := h.CustomerService.GetByUsername(c, username)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(customer))
}

func (h handlerCustomer) GetByEmail(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()
	email := ctx.Params("email")
	customer, err := h.CustomerService.GetByEmail(c, email)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(customer))
}

func (h handlerCustomer) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	customers, err := h.CustomerService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(customers))
}

func (h handlerCustomer) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	id := ctx.Params("id")

	err := h.CustomerService.Delete(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(response.ResponseError(err.Error(), http.StatusInternalServerError))
	}

	return ctx.Status(http.StatusOK).JSON(response.ResponseSuccess(""))
}
