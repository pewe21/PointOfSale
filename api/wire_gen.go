// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package api

import (
	"database/sql"
	"github.com/pewe21/PointOfSale/internal/config"
	"github.com/pewe21/PointOfSale/internal/domain"
	"github.com/pewe21/PointOfSale/internal/handler"
	"github.com/pewe21/PointOfSale/internal/module/authentication"
	"github.com/pewe21/PointOfSale/internal/module/brand"
	"github.com/pewe21/PointOfSale/internal/module/customer"
	"github.com/pewe21/PointOfSale/internal/module/customer_roles"
	"github.com/pewe21/PointOfSale/internal/module/product"
	"github.com/pewe21/PointOfSale/internal/module/role"
	"github.com/pewe21/PointOfSale/internal/module/supplier"
	"github.com/pewe21/PointOfSale/internal/module/user"
)

// Injectors from injector.go:

func InitializedUser(conn *sql.DB) domain.UserHandler {
	userRepository := user.NewRepository(conn)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	return userHandler
}

func InitializedAuthentication(conn *sql.DB, cnf *config.Jwt) domain.AuthHandler {
	userRepository := user.NewRepository(conn)
	authService := authentication.NewService(userRepository, cnf)
	authHandler := handler.NewAuthenticationHandler(authService)
	return authHandler
}

func InitializedCustomer(conn *sql.DB) domain.CustomerHandler {
	customerRepository := customer.NewRepository(conn)
	roleRepository := role.NewRepository(conn)
	customerRolesRepository := customer_roles.NewCustomerRolesRepository(conn)
	customerService := customer.NewService(customerRepository, roleRepository, customerRolesRepository)
	customerHandler := handler.NewHandlerCustomer(customerService)
	return customerHandler
}

func InitializedSupplier(conn *sql.DB) domain.SupplierHandler {
	supplierRepository := supplier.NewRepository(conn)
	supplierService := supplier.NewService(supplierRepository)
	supplierHandler := handler.NewHandlerSupplier(supplierService)
	return supplierHandler
}

func InitializedBrand(conn *sql.DB) domain.BrandHandler {
	brandRepository := brand.NewRepository(conn)
	brandService := brand.NewService(brandRepository)
	brandHandler := handler.NewHandlerType(brandService)
	return brandHandler
}

func InitializedProduct(conn *sql.DB) domain.ProductHandler {
	productRepository := product.NewRepository(conn)
	productService := product.NewService(productRepository)
	productHandler := handler.NewHandlerProduct(productService)
	return productHandler
}

func InitializedRole(conn *sql.DB) domain.RoleHandler {
	roleRepository := role.NewRepository(conn)
	roleService := role.NewService(roleRepository)
	roleHandler := handler.NewHandlerRole(roleService)
	return roleHandler
}
