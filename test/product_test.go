package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/PointOfSale/api"
	"github.com/pewe21/PointOfSale/dto"
	"github.com/pewe21/PointOfSale/internal/domain"
	"github.com/stretchr/testify/assert"
)

func deleteAllProduct(conn *sql.DB) {
	conn.Exec("DELETE FROM products")
}

func deleteBrandTestProduct(conn *sql.DB) {
	conn.Exec("DELETE FROM brands where name = 'Brand 1'")
}

func deleteSupplierTestProduct(conn *sql.DB) {
	conn.Exec("DELETE FROM suppliers where name = 'Supplier 1'")
}

func ProductSetup() *fiber.App {
	conn := GlobalSetupTest()
	deleteAllProduct(conn)
	deleteBrandTestProduct(conn)
	deleteSupplierTestProduct(conn)
	app := fiber.New()
	api.NewProductApi(app, conn)
	return app
}

func GlobalCreateProduct(t *testing.T, app *fiber.App, product dto.CreateProductRequest) *http.Response {
	body, _ := json.Marshal(product)

	req := httptest.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	return resp
}

func GlobalGetProduct(t *testing.T, app *fiber.App) *http.Response {
	req := httptest.NewRequest(http.MethodGet, "/product", nil)
	resp, _ := app.Test(req)

	return resp
}

func TestCreateProduct(t *testing.T) {
	var BrandId string
	var SupplierId string
	appBrand := BrandSetup()
	appSupplier := SupplierSetup()
	app := ProductSetup()

	t.Run("Create Brand", func(t *testing.T) {
		brand := dto.CreateBrandRequest{
			Name:        "Brand 1",
			Description: "Brand 1 Description",
		}

		resp := GlobalCreateBrand(t, appBrand, brand)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})
	t.Run("Create Supplier", func(t *testing.T) {
		supplier := dto.CreateSupplierRequest{
			Name:    "Supplier 1",
			Email:   "romadhon@emal.com",
			Address: "Jl. Jalan",
			Phone:   "08123456789",
		}

		resp := GlobalCreateSupplier(t, appSupplier, supplier)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})
	t.Run("Get Brand and Supplier After Create", func(t *testing.T) {
		resp := GlobalGetBrand(t, appBrand)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var fetchedBrand FetchedResponse[[]domain.Brand]
		json.NewDecoder(resp.Body).Decode(&fetchedBrand)

		BrandId = fetchedBrand.Data[0].Id

		resp = GlobalGetSupplier(t, appSupplier)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var fetchedSupplier FetchedResponse[[]domain.Supplier]
		json.NewDecoder(resp.Body).Decode(&fetchedSupplier)

		SupplierId = fetchedSupplier.Data[0].Id

	})

	t.Run("Create Product", func(t *testing.T) {
		log.Println("ID Brand: ", BrandId)

		product := dto.CreateProductRequest{
			Name:       "Product 1",
			BrandId:    BrandId,
			SupplierId: SupplierId,
			SKU:        "ttt",
		}

		resp := GlobalCreateProduct(t, app, product)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		var createdProduct FetchedResponse[domain.Product]
		json.NewDecoder(resp.Body).Decode(&createdProduct)
		assert.Equal(t, http.StatusCreated, createdProduct.Code)
		assert.Empty(t, createdProduct.Data)
	})
}
