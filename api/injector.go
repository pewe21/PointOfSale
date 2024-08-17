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
