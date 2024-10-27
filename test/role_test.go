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
	"github.com/pewe21/PointOfSale/internal/domain"
	"github.com/stretchr/testify/assert"
)

// run this test with command "go test ./test/setup_test.go ./test/role_test.go -v"

func deleteAllRole(conn *sql.DB) {
	conn.Exec("DELETE FROM roles")
}

func setup() (*fiber.App, *sql.DB) {
	conn := GlobalSetupTest()
	app := fiber.New()
	api.NewRoleApi(app, conn)
	return app, conn
}

func TestCreateRole(t *testing.T) {
	app, conn := setup()
	deleteAllRole(conn)
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

func TestCreateDuplicateRole(t *testing.T) {
	app, conn := setup()
	deleteAllRole(conn)

	// create 1st role
	role := dto.CreateRoleRequest{Name: "Admin", DisplayName: "Administrator"}
	body, _ := json.Marshal(role)

	req := httptest.NewRequest(http.MethodPost, "/role", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	// create 2nd role
	body, _ = json.Marshal(role)

	req = httptest.NewRequest(http.MethodPost, "/role", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = app.Test(req)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

	var createdRole struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    domain.Role `json:"data"`
	}
	json.NewDecoder(resp.Body).Decode(&createdRole)
	assert.Equal(t, "role already exist", createdRole.Message)
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

func TestDeleteRole(t *testing.T) {
	app, _ := setup()
	//get the role
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

	// Now delete the role
	req = httptest.NewRequest(http.MethodDelete, "/role/"+fetchedRole.Data[0].Id, nil)
	resp, _ = app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var newFetchedRole struct {
		Code    int            `json:"code"`
		Message string         `json:"message"`
		Data    []dto.RoleData `json:"data"`
	}
	json.NewDecoder(resp.Body).Decode(&newFetchedRole)
	// Verify the role is deleted
	req = httptest.NewRequest(http.MethodGet, "/role/"+fetchedRole.Data[0].Id, nil)
	resp, _ = app.Test(req)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Empty(t, newFetchedRole.Data)
}

func TestDeleteWrongIdRole(t *testing.T) {
	app, _ := setup()

	// Now delete the role
	req := httptest.NewRequest(http.MethodDelete, "/role/"+"asadasd", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	var newFetchedRole struct {
		Code    int            `json:"code"`
		Message string         `json:"message"`
		Data    []dto.RoleData `json:"data"`
	}
	json.NewDecoder(resp.Body).Decode(&newFetchedRole)
	assert.Equal(t, newFetchedRole.Message, "role not found")
	assert.Empty(t, newFetchedRole.Data)
}
