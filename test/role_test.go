package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/PointOfSale/api"
	"github.com/pewe21/PointOfSale/dto"
	"github.com/pewe21/PointOfSale/internal/config"
	"github.com/pewe21/PointOfSale/internal/database"
	"github.com/pewe21/PointOfSale/internal/domain"
	"github.com/stretchr/testify/assert"
)

func InitializedLoader() *config.Config {

	return &config.Config{
		Server: config.Server{
			Host: "127.0.0.1",
			Port: "3000",
		},
		Database: config.Database{
			Host:     "localhost",
			Port:     "5432",
			User:     "pos",
			Password: "passpos",
			Name:     "test_pointOfSale",
			Tz:       "Asia/Jakarta",
		},
		Jwt: config.Jwt{
			Secret: "secret",
			Exp:    60,
		},
	}
}

func DeleteAllRole(conn *sql.DB) {
	conn.Exec("DELETE FROM roles")
}

func setup() (*fiber.App, *sql.DB) {
	conf := InitializedLoader()

	conn := database.InitDB(conf.Database, false)
	app := fiber.New()
	api.NewRoleApi(app, conn)
	return app, conn
}

func TestCreateRole(t *testing.T) {
	app, conn := setup()
	DeleteAllRole(conn)
	role := dto.CreateRoleRequest{Name: "Admin", DisplayName: "Administrator"}
	body, _ := json.Marshal(role)

	req := httptest.NewRequest(http.MethodPost, "/role", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var createdRole struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    domain.Role `json:"data"`
	}
	json.NewDecoder(resp.Body).Decode(&createdRole)
	assert.Equal(t, http.StatusCreated, createdRole.Code)
}

func TestGetRole(t *testing.T) {
	app, _ := setup()

	// Now get the role
	req := httptest.NewRequest(http.MethodGet, "/role", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var fetchedRole struct {
		Code    int            `json:"code"`
		Message string         `json:"message"`
		Data    []dto.RoleData `json:"data"`
	}
	json.NewDecoder(resp.Body).Decode(&fetchedRole)
	assert.Equal(t, "Admin", fetchedRole.Data[0].Name)
}

// func TestUpdateRole(t *testing.T) {
// 	app := setup()
// 	role := Role{Name: "Editor"}
// 	body, _ := json.Marshal(role)

// 	req := httptest.NewRequest(http.MethodPost, "/roles", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	var createdRole Role
// 	json.NewDecoder(resp.Body).Decode(&createdRole)

// 	// Update the role
// 	updatedRole := Role{Name: "Super Editor"}
// 	body, _ = json.Marshal(updatedRole)

// 	req = httptest.NewRequest(http.MethodPut, "/roles/"+createdRole.ID, bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ = app.Test(req)

// 	assert.Equal(t, http.StatusOK, resp.StatusCode)

// 	var updatedRoleResponse Role
// 	json.NewDecoder(resp.Body).Decode(&updatedRoleResponse)
// 	assert.Equal(t, "Super Editor", updatedRoleResponse.Name)
// }

// func TestDeleteRole(t *testing.T) {
// 	app := setup()
// 	role := Role{Name: "Guest"}
// 	body, _ := json.Marshal(role)

// 	// Create a role first
// 	req := httptest.NewRequest(http.MethodPost, "/roles", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req)

// 	var createdRole Role
// 	json.NewDecoder(resp.Body).Decode(&createdRole)

// 	// Now delete the role
// 	req = httptest.NewRequest(http.MethodDelete, "/roles/"+createdRole.ID, nil)
// 	resp, _ = app.Test(req)

// 	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

// 	// Verify the role is deleted
// 	req = httptest.NewRequest(http.MethodGet, "/roles/"+createdRole.ID, nil)
// 	resp, _ = app.Test(req)

// 	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
// }
