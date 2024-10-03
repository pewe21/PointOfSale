//go:build wireinject
// +build wireinject

package api

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/pewe21/PointOfSale/internal/config"
	"github.com/pewe21/PointOfSale/internal/domain"
	"github.com/pewe21/PointOfSale/internal/handler"
	"github.com/pewe21/PointOfSale/internal/module/authentication"
	"github.com/pewe21/PointOfSale/internal/module/brand"
	"github.com/pewe21/PointOfSale/internal/module/customer"
	"github.com/pewe21/PointOfSale/internal/module/product"
	"github.com/pewe21/PointOfSale/internal/module/supplier"
	"github.com/pewe21/PointOfSale/internal/module/user"
)

func InitializedUser(conn *sql.DB) domain.UserHandler {
	wire.Build(user.NewRepository, user.NewService, handler.NewUserHandler)
	return nil
}

func InitializedAuthentication(conn *sql.DB, cnf *config.Jwt) domain.AuthHandler {
	wire.Build(user.NewRepository, authentication.NewService, handler.NewAuthenticationHandler)
	return nil
}

func InitializedCustomer(conn *sql.DB) domain.CustomerHandler {
	wire.Build(customer.NewRepository, customer.NewService, handler.NewHandlerCustomer)
	return nil
}

func InitializedSupplier(conn *sql.DB) domain.SupplierHandler {
	wire.Build(supplier.NewRepository, supplier.NewService, handler.NewHandlerSupplier)
	return nil
}

func InitializedBrand(conn *sql.DB) domain.BrandHandler {
	wire.Build(brand.NewRepository, brand.NewService, handler.NewHandlerType)
	return nil
}

func InitializedProduct(conn *sql.DB) domain.ProductHandler {
	wire.Build(product.NewRepository, product.NewService, handler.NewHandlerProduct)
	return nil
}
