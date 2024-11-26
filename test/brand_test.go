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
	app := fiber.New()
	conn := GlobalSetupTest()
	deleteAllBrand(conn)
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
	t.Run("Create Brand Success", func(t *testing.T) {
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
	t.Run("Create Brand Failed", func(t *testing.T) {
		brand := dto.CreateBrandRequest{
			Name:        "",
			Description: "Brand failed Description",
		}

		resp := GlobalCreateBrand(t, app, brand)
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
		}
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var createdBrand FetchedResponse[domain.Brand]
		json.NewDecoder(resp.Body).Decode(&createdBrand)
		assert.Equal(t, http.StatusBadRequest, createdBrand.Code)
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
		assert.Equal(t, http.StatusConflict, resp.StatusCode)
		assert.Equal(t, "brand already exist", fetchedBrand.Message)
	})
}

func TestUpdateBrand(t *testing.T) {
	app := BrandSetup()

	brand := dto.CreateBrandRequest{
		Name:        "Brand 1",
		Description: "Brand 1 Description",
	}

	resp := GlobalCreateBrand(t, app, brand)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	getBrand := GlobalGetBrand(t, app)
	assert.Equal(t, http.StatusOK, getBrand.StatusCode)

	var getBrandData FetchedResponse[[]dto.BrandData]
	json.NewDecoder(getBrand.Body).Decode(&getBrandData)
	assert.NotEmpty(t, getBrandData.Data)

	t.Run("Update Brand Success", func(t *testing.T) {
		update := dto.UpdateBrandRequest{
			Name:        "Brand 2",
			Description: "Brand 2 Description",
		}

		body, _ := json.Marshal(update)

		req := httptest.NewRequest(http.MethodPut, "/brand/"+getBrandData.Data[0].Id, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		res, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
		}

		var updatedBrand FetchedResponse[domain.Brand]

		json.NewDecoder(res.Body).Decode(&updatedBrand)

		newReq := httptest.NewRequest(http.MethodGet, "/brand/"+getBrandData.Data[0].Id, nil)

		newResp, _ := app.Test(newReq)

		if newResp.StatusCode != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, newResp.StatusCode)
		}

		json.NewDecoder(newResp.Body).Decode(&updatedBrand)

		assert.Equal(t, update.Name, updatedBrand.Data.Name)

	})

	t.Run("Update Brand Failed", func(t *testing.T) {
		update := dto.UpdateBrandRequest{
			Name:        "",
			Description: "Brand 2 Description",
		}

		body, _ := json.Marshal(update)

		req := httptest.NewRequest(http.MethodPut, "/brand/"+getBrandData.Data[0].Id, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		res, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, res.StatusCode)
		}

		var updatedBrand FetchedResponse[domain.Brand]

		json.NewDecoder(res.Body).Decode(&updatedBrand)

		newReq := httptest.NewRequest(http.MethodGet, "/brand/"+getBrandData.Data[0].Id, nil)

		newResp, _ := app.Test(newReq)

		if newResp.StatusCode != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, newResp.StatusCode)
		}
		json.NewDecoder(newResp.Body).Decode(&updatedBrand)

		assert.NotEmpty(t, updatedBrand.Data.Name)
		assert.NotSame(t, update.Name, updatedBrand.Data.Name)
	})
}
