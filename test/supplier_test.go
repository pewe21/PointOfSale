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
	"github.com/pewe21/PointOfSale/internal/response"
	"github.com/stretchr/testify/assert"
)

// run this test with command "go test ./test/setup_test.go ./test/role_test.go -v"

func deleteAllSupplier(conn *sql.DB) {
	conn.Exec("DELETE FROM suppliers")
}

func SupplierSetup() *fiber.App {
	conn := GlobalSetupTest()
	deleteAllSupplier(conn)
	app := fiber.New()
	api.NewSupplierApi(app, conn)
	return app
}

func GlobalCreateSupplier(t *testing.T, app *fiber.App, supplier dto.CreateSupplierRequest) *http.Response {
	body, _ := json.Marshal(supplier)

	req := httptest.NewRequest(http.MethodPost, "/supplier", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	return resp
}

func GlobalGetSupplierRole(t *testing.T, app *fiber.App) *http.Response {
	req := httptest.NewRequest(http.MethodGet, "/supplier", nil)
	resp, _ := app.Test(req)

	return resp
}

func TestCreateSupplier(t *testing.T) {
	app := SupplierSetup()
	t.Run("Create Supplier", func(t *testing.T) {
		supplier := dto.CreateSupplierRequest{
			Name:    "Supplier 1",
			Email:   "romadhon@emal.com",
			Address: "Jl. Jalan",
			Phone:   "08123456789",
		}

		resp := GlobalCreateSupplier(t, app, supplier)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		var createdSupplier struct {
			Code    int             `json:"code"`
			Message string          `json:"message"`
			Data    domain.Supplier `json:"data"`
		}
		json.NewDecoder(resp.Body).Decode(&createdSupplier)
		assert.Equal(t, http.StatusCreated, createdSupplier.Code)
		assert.Empty(t, createdSupplier.Data)
	})

}

func TestCreateDuplicateSupplier(t *testing.T) {
	app := SupplierSetup()

	// 	//one data
	supplier := dto.CreateSupplierRequest{
		Name:    "Supplier 1",
		Email:   "romadhon@emal.com",
		Address: "Jl. Jalan",
		Phone:   "08123456789",
	}
	// create 1st supplier
	t.Run("Create Supplier 1", func(t *testing.T) {
		resp := GlobalCreateSupplier(t, app, supplier)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})

	// create 2nd role
	t.Run("Create Supplier 2", func(t *testing.T) {
		resp := GlobalCreateSupplier(t, app, supplier)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		var createdSupplier response.Response[string]
		json.NewDecoder(resp.Body).Decode(&createdSupplier)
		assert.Equal(t, "supplier already exist", createdSupplier.Message)
		assert.Empty(t, createdSupplier.Data)
	})
}

// func TestGetRole(t *testing.T) {
// 	app := setup()

// 	role := dto.CreateRoleRequest{
// 		Name:        "Admin",
// 		DisplayName: "Administator",
// 	}
// 	t.Run("Create Role", func(t *testing.T) {
// 		resp := GlobalCreateRole(t, app, role)
// 		assert.Equal(t, http.StatusCreated, resp.StatusCode)
// 	})

// 	// Now get the role
// 	t.Run("Get Role", func(t *testing.T) {
// 		resp := GlobalGetRole(t, app)
// 		assert.Equal(t, http.StatusOK, resp.StatusCode)

// 		var fetchedRole struct {
// 			Code    int            `json:"code"`
// 			Message string         `json:"message"`
// 			Data    []dto.RoleData `json:"data"`
// 		}
// 		json.NewDecoder(resp.Body).Decode(&fetchedRole)
// 		assert.NotEmpty(t, fetchedRole.Data)
// 		assert.Equal(t, "Admin", fetchedRole.Data[0].Name)
// 	})
// }

// func TestUpdateRole(t *testing.T) {
// 	app := setup()

// 	//create role
// 	role := dto.CreateRoleRequest{
// 		Name:        "Admin",
// 		DisplayName: "Administator",
// 	}

// 	resp := GlobalCreateRole(t, app, role)
// 	assert.Equal(t, http.StatusCreated, resp.StatusCode)

// 	// Now get the role to get the id
// 	resp = GlobalGetRole(t, app)
// 	assert.Equal(t, http.StatusOK, resp.StatusCode)

// 	var fetchedRole struct {
// 		Code    int            `json:"code"`
// 		Message string         `json:"message"`
// 		Data    []dto.RoleData `json:"data"`
// 	}
// 	json.NewDecoder(resp.Body).Decode(&fetchedRole)
// 	assert.Equal(t, "Administator", fetchedRole.Data[0].DisplayName)

// 	// Now update the role
// 	updatedRole := dto.UpdateRoleRequest{DisplayName: "Super Editor"}
// 	body, _ := json.Marshal(updatedRole)

// 	req := httptest.NewRequest(http.MethodPut, "/role/"+fetchedRole.Data[0].Id, bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ = app.Test(req)

// 	assert.Equal(t, http.StatusOK, resp.StatusCode)

// 	json.NewDecoder(resp.Body).Decode(&fetchedRole)
// 	assert.Equal(t, "success", fetchedRole.Message)

// 	// Now get the role again
// 	resp = GlobalGetRole(t, app)
// 	assert.Equal(t, http.StatusOK, resp.StatusCode)

// 	json.NewDecoder(resp.Body).Decode(&fetchedRole)
// 	assert.NotEmpty(t, fetchedRole.Data)
// 	assert.NotEqual(t, "", fetchedRole.Data[0].Id)
// 	assert.Equal(t, "Admin", fetchedRole.Data[0].Name)
// 	assert.Equal(t, "Super Editor", fetchedRole.Data[0].DisplayName)

// }

// func TestDeleteRole(t *testing.T) {
// 	app := setup()
// 	role := dto.CreateRoleRequest{
// 		Name:        "Admin",
// 		DisplayName: "Administator",
// 	}

// 	resp := GlobalCreateRole(t, app, role)
// 	assert.Equal(t, http.StatusCreated, resp.StatusCode)

// 	// Now get the role
// 	resp = GlobalGetRole(t, app)
// 	assert.Equal(t, http.StatusOK, resp.StatusCode)

// 	var fetchedRole struct {
// 		Code    int            `json:"code"`
// 		Message string         `json:"message"`
// 		Data    []dto.RoleData `json:"data"`
// 	}
// 	json.NewDecoder(resp.Body).Decode(&fetchedRole)
// 	assert.NotEqual(t, "", fetchedRole.Data[0].Id)
// 	assert.Equal(t, "Admin", fetchedRole.Data[0].Name)

// 	// Now delete the role
// 	req := httptest.NewRequest(http.MethodDelete, "/role/"+fetchedRole.Data[0].Id, nil)
// 	resp, _ = app.Test(req)

// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// 	var newFetchedRole struct {
// 		Code    int            `json:"code"`
// 		Message string         `json:"message"`
// 		Data    []dto.RoleData `json:"data"`
// 	}
// 	json.NewDecoder(resp.Body).Decode(&newFetchedRole)
// 	// Verify the role is deleted
// 	req = httptest.NewRequest(http.MethodGet, "/role/"+fetchedRole.Data[0].Id, nil)
// 	resp, _ = app.Test(req)

// 	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
// 	assert.Empty(t, newFetchedRole.Data)
// }

// func TestDeleteWrongIdRole(t *testing.T) {
// 	app := setup()

// 	// Now delete the role
// 	req := httptest.NewRequest(http.MethodDelete, "/role/"+"asadasd", nil)
// 	resp, _ := app.Test(req)

// 	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
// 	var newFetchedRole struct {
// 		Code    int            `json:"code"`
// 		Message string         `json:"message"`
// 		Data    []dto.RoleData `json:"data"`
// 	}
// 	json.NewDecoder(resp.Body).Decode(&newFetchedRole)
// 	assert.Equal(t, newFetchedRole.Message, "role not found")
// 	assert.Empty(t, newFetchedRole.Data)
// }
