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

func deleteAllBrand(conn *sql.DB) {
	conn.Exec("DELETE FROM products") // to avoid foreign key constraint
	conn.Exec("DELETE FROM brands")
}

func BrandSetup() *fiber.App {
	conn := GlobalSetupTest()
	deleteAllBrand(conn)
	app := fiber.New()
	api.NewBrandApi(app, conn)
	return app
}

func GlobalCreateBrand(t *testing.T, app *fiber.App, role dto.CreateBrandRequest) *http.Response {
	body, _ := json.Marshal(role)

	req := httptest.NewRequest(http.MethodPost, "/brand", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	return resp
}

func GlobalGetBrand(t *testing.T, app *fiber.App) *http.Response {
	req := httptest.NewRequest(http.MethodGet, "/brand", nil)
	resp, _ := app.Test(req)

	return resp
}

func TestCreateBrand(t *testing.T) {
	app := BrandSetup()
	t.Run("Create Brand", func(t *testing.T) {
		brand := dto.CreateBrandRequest{
			Name:        "Brand 1",
			Description: "Brand 1 Description",
		}

		resp := GlobalCreateBrand(t, app, brand)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		var createdBrand FetchedResponse[domain.Brand]
		json.NewDecoder(resp.Body).Decode(&createdBrand)
		assert.Equal(t, http.StatusCreated, createdBrand.Code)
		assert.Empty(t, createdBrand.Data)
	})
}

func TestCreateDuplicateBrand(t *testing.T) {
	app := BrandSetup()
	brand := dto.CreateBrandRequest{
		Name:        "Brand 1",
		Description: "Brand 1 Description",
	}

	t.Run("Create Brand 1", func(t *testing.T) {
		resp := GlobalCreateBrand(t, app, brand)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

	})

	var fetchedBrand FetchedResponse[domain.Brand]

	t.Run("Create Brand 2", func(t *testing.T) {
		resp := GlobalCreateBrand(t, app, brand)

		json.NewDecoder(resp.Body).Decode(&fetchedBrand)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		assert.Equal(t, "error saving brand, brand already exist", fetchedBrand.Message)
	})
}
