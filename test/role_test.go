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

func RoleSetup() *fiber.App {
	conn := GlobalSetupTest()
	deleteAllRole(conn)
	app := fiber.New()
	api.NewRoleApi(app, conn)
	return app
}

func GlobalCreateRole(t *testing.T, app *fiber.App, role dto.CreateRoleRequest) *http.Response {
	body, _ := json.Marshal(role)

	req := httptest.NewRequest(http.MethodPost, "/role", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	return resp
}

func GlobalGetRole(t *testing.T, app *fiber.App) *http.Response {
	req := httptest.NewRequest(http.MethodGet, "/role", nil)
	resp, _ := app.Test(req)

	return resp
}

func TestCreateRole(t *testing.T) {
	app := RoleSetup()
	t.Run("Create Role", func(t *testing.T) {
		role := dto.CreateRoleRequest{
			Name:        "Admin",
			DisplayName: "Administator",
		}

		resp := GlobalCreateRole(t, app, role)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		var createdRole FetchedResponse[domain.Role]
		json.NewDecoder(resp.Body).Decode(&createdRole)
		assert.Equal(t, http.StatusCreated, createdRole.Code)
		assert.Empty(t, createdRole.Data)
	})

}

func TestCreateDuplicateRole(t *testing.T) {
	app := RoleSetup()

	//one data
	role := dto.CreateRoleRequest{
		Name:        "Admin",
		DisplayName: "Administator",
	}
	// create 1st role
	t.Run("Create Role 1", func(t *testing.T) {
		resp := GlobalCreateRole(t, app, role)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})

	// create 2nd role
	t.Run("Create Role 2", func(t *testing.T) {
		resp := GlobalCreateRole(t, app, role)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		var createdRole FetchedResponse[domain.Role]
		json.NewDecoder(resp.Body).Decode(&createdRole)
		assert.Equal(t, "role already exist", createdRole.Message)
		assert.Empty(t, createdRole.Data)
	})
}

func TestGetRole(t *testing.T) {
	app := RoleSetup()

	role := dto.CreateRoleRequest{
		Name:        "Admin",
		DisplayName: "Administator",
	}
	t.Run("Create Role", func(t *testing.T) {
		resp := GlobalCreateRole(t, app, role)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})

	// Now get the role
	t.Run("Get Role", func(t *testing.T) {
		resp := GlobalGetRole(t, app)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var fetchedRole FetchedResponse[[]domain.Role]
		json.NewDecoder(resp.Body).Decode(&fetchedRole)
		assert.NotEmpty(t, fetchedRole.Data)
		assert.Equal(t, "Admin", fetchedRole.Data[0].Name)
	})
}

func TestUpdateRole(t *testing.T) {
	app := RoleSetup()

	//create role
	role := dto.CreateRoleRequest{
		Name:        "Admin",
		DisplayName: "Administator",
	}

	resp := GlobalCreateRole(t, app, role)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Now get the role to get the id
	resp = GlobalGetRole(t, app)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var fetchedRole FetchedResponse[[]domain.Role]
	json.NewDecoder(resp.Body).Decode(&fetchedRole)
	assert.Equal(t, "Administator", fetchedRole.Data[0].DisplayName)

	// Now update the role
	updatedRole := dto.UpdateRoleRequest{DisplayName: "Super Editor"}
	body, _ := json.Marshal(updatedRole)

	req := httptest.NewRequest(http.MethodPut, "/role/"+fetchedRole.Data[0].Id, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	json.NewDecoder(resp.Body).Decode(&fetchedRole)
	assert.Equal(t, "success", fetchedRole.Message)

	// Now get the role again
	resp = GlobalGetRole(t, app)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	json.NewDecoder(resp.Body).Decode(&fetchedRole)
	assert.NotEmpty(t, fetchedRole.Data)
	assert.NotEqual(t, "", fetchedRole.Data[0].Id)
	assert.Equal(t, "Admin", fetchedRole.Data[0].Name)
	assert.Equal(t, "Super Editor", fetchedRole.Data[0].DisplayName)

}

func TestDeleteRole(t *testing.T) {
	app := RoleSetup()
	role := dto.CreateRoleRequest{
		Name:        "Admin",
		DisplayName: "Administator",
	}

	resp := GlobalCreateRole(t, app, role)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Now get the role
	resp = GlobalGetRole(t, app)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var fetchedRole FetchedResponse[[]domain.Role]
	json.NewDecoder(resp.Body).Decode(&fetchedRole)
	assert.NotEqual(t, "", fetchedRole.Data[0].Id)
	assert.Equal(t, "Admin", fetchedRole.Data[0].Name)

	// Now delete the role
	req := httptest.NewRequest(http.MethodDelete, "/role/"+fetchedRole.Data[0].Id, nil)
	resp, _ = app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var newFetchedRole FetchedResponse[[]domain.Role]
	json.NewDecoder(resp.Body).Decode(&newFetchedRole)
	// Verify the role is deleted
	req = httptest.NewRequest(http.MethodGet, "/role/"+fetchedRole.Data[0].Id, nil)
	resp, _ = app.Test(req)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Empty(t, newFetchedRole.Data)
}

func TestDeleteWrongIdRole(t *testing.T) {
	app := RoleSetup()

	// Now delete the role
	req := httptest.NewRequest(http.MethodDelete, "/role/"+"asadasd", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	var newFetchedRole FetchedResponse[[]domain.Role]
	json.NewDecoder(resp.Body).Decode(&newFetchedRole)
	assert.Equal(t, newFetchedRole.Message, "role not found")
	assert.Empty(t, newFetchedRole.Data)
}
