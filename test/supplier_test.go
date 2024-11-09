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

func GlobalGetSupplier(t *testing.T, app *fiber.App) *http.Response {
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

		var createdSupplier FetchedResponse[domain.Supplier]

		json.NewDecoder(resp.Body).Decode(&createdSupplier)
		assert.Equal(t, http.StatusCreated, createdSupplier.Code)
		assert.Empty(t, createdSupplier.Data)
	})

}

func TestCreateDuplicateSupplier(t *testing.T) {
	app := SupplierSetup()

	// data
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

func TestGetSupplier(t *testing.T) {
	app := SupplierSetup()

	supplier := dto.CreateSupplierRequest{
		Name:    "Supplier 1",
		Email:   "romadhon@emal.com",
		Address: "Jl. Jalan",
		Phone:   "08123456789",
	}
	t.Run("Create Supplier", func(t *testing.T) {
		resp := GlobalCreateSupplier(t, app, supplier)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})

	// Now get the supplier
	t.Run("Get Supplier", func(t *testing.T) {
		resp := GlobalGetSupplier(t, app)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var fetchedSupplier FetchedResponse[[]domain.Supplier]
		json.NewDecoder(resp.Body).Decode(&fetchedSupplier)
		assert.NotEmpty(t, fetchedSupplier.Data)
		assert.Equal(t, "Supplier 1", fetchedSupplier.Data[0].Name)
		assert.Equal(t, "romadhon@emal.com", fetchedSupplier.Data[0].Email)
		assert.Equal(t, "Jl. Jalan", fetchedSupplier.Data[0].Address)
		assert.Equal(t, "08123456789", fetchedSupplier.Data[0].Phone)
	})
}

func TestUpdateSupplier(t *testing.T) {
	app := SupplierSetup()

	//create supplier
	supplier := dto.CreateSupplierRequest{
		Name:    "Supplier 1",
		Email:   "romadhon@emal.com",
		Address: "Jl. Jalan",
		Phone:   "08123456789",
	}

	resp := GlobalCreateSupplier(t, app, supplier)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Now get the supplier to get the id
	resp = GlobalGetSupplier(t, app)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var fetchedSupplier FetchedResponse[[]domain.Supplier]
	json.NewDecoder(resp.Body).Decode(&fetchedSupplier)
	assert.Equal(t, "Supplier 1", fetchedSupplier.Data[0].Name)

	// Now update the supplier
	updatedSupplier := dto.UpdateSupplierRequest{
		Name:    "Supplier 1 Ubah",
		Email:   "ubah@emal.com",
		Address: "Jl. Jalan Ke Kota",
		Phone:   "08123456789",
	}
	body, _ := json.Marshal(updatedSupplier)

	newUri := "/supplier/" + fetchedSupplier.Data[0].Id

	req := httptest.NewRequest(http.MethodPut, newUri, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	json.NewDecoder(resp.Body).Decode(&fetchedSupplier)
	assert.Equal(t, "success", fetchedSupplier.Message)

	// Now get the supplier again
	resp = GlobalGetSupplier(t, app)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	json.NewDecoder(resp.Body).Decode(&fetchedSupplier)
	assert.NotEmpty(t, fetchedSupplier.Data)
	assert.NotEqual(t, "", fetchedSupplier.Data[0].Id)
	assert.Equal(t, "Supplier 1 Ubah", fetchedSupplier.Data[0].Name)
	assert.Equal(t, "ubah@emal.com", fetchedSupplier.Data[0].Email)
	assert.Equal(t, "Jl. Jalan Ke Kota", fetchedSupplier.Data[0].Address)
	assert.Equal(t, "08123456789", fetchedSupplier.Data[0].Phone)

}

// func TestDeleteRole(t *testing.T) {
// 	app := SupplierSetup()
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
// 	app := SupplierSetup()

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
