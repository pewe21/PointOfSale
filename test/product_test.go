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
	conn.Exec("DELETE FROM brands")
}

func deleteSupplierTestProduct(conn *sql.DB) {
	conn.Exec("DELETE FROM suppliers")
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

func GlobalGetProductId(t *testing.T, app *fiber.App, id string) *http.Response {
	req := httptest.NewRequest(http.MethodGet, "/product/"+id, nil)
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

func TestDuplicateProductSKU(t *testing.T) {
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
		product := dto.CreateProductRequest{
			Name:       "Product 1",
			BrandId:    BrandId,
			SupplierId: SupplierId,
			SKU:        "ttt",
		}
		t.Run("Create Product 1", func(t *testing.T) {
			resp := GlobalCreateProduct(t, app, product)
			assert.Equal(t, http.StatusCreated, resp.StatusCode)
		})

		t.Run("Create Duplicate Product SKU", func(t *testing.T) {
			resp := GlobalCreateProduct(t, app, product)
			assert.Equal(t, http.StatusConflict, resp.StatusCode)

			var createdProduct FetchedResponse[domain.Product]
			json.NewDecoder(resp.Body).Decode(&createdProduct)
			assert.Equal(t, http.StatusConflict, createdProduct.Code)
			assert.Equal(t, "cannot create product, SKU already exist", createdProduct.Message)
			assert.Empty(t, createdProduct.Data)
		})

	})
}

func TestUpdateProduct(t *testing.T) {
	var BrandId string
	var SupplierId string

	var ProductId string
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
		product := dto.CreateProductRequest{
			Name:       "Product 1",
			BrandId:    BrandId,
			SupplierId: SupplierId,
			SKU:        "ttt",
		}
		resp := GlobalCreateProduct(t, app, product)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

	})

	t.Run("Get Product", func(t *testing.T) {
		resp := GlobalGetProduct(t, app)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var fetchedProduct FetchedResponse[[]domain.Product]
		json.NewDecoder(resp.Body).Decode(&fetchedProduct)
		assert.NotEmpty(t, fetchedProduct.Data)

		ProductId = fetchedProduct.Data[0].Id
	})

	t.Run("Update Product", func(t *testing.T) {
		product := dto.UpdateProductRequest{
			Name:       "Product Sudah Diubah",
			SKU:        "ttt",
			BrandId:    BrandId,
			SupplierId: SupplierId,
		}
		body, _ := json.Marshal(product)

		req := httptest.NewRequest(http.MethodPut, "/product/"+ProductId, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Get Product After Update", func(t *testing.T) {
		resp := GlobalGetProductId(t, app, ProductId)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var fetchedProduct FetchedResponse[domain.Product]
		json.NewDecoder(resp.Body).Decode(&fetchedProduct)
		assert.NotEmpty(t, fetchedProduct.Data)
		assert.Equal(t, "Product Sudah Diubah", fetchedProduct.Data.Name)
	})
}

func TestDeleteProduct(t *testing.T) {
	var BrandId string
	var SupplierId string

	var ProductId string
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
			Email:   "Supplier@s.com",
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
		product := dto.CreateProductRequest{
			Name:       "Product 1",
			BrandId:    BrandId,
			SupplierId: SupplierId,
			SKU:        "ttt",
		}
		resp := GlobalCreateProduct(t, app, product)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

	})

	t.Run("Get Product", func(t *testing.T) {

		resp := GlobalGetProduct(t, app)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var fetchedProduct FetchedResponse[[]domain.Product]
		json.NewDecoder(resp.Body).Decode(&fetchedProduct)

		assert.NotEmpty(t, fetchedProduct.Data)
		ProductId = fetchedProduct.Data[0].Id

	})

	t.Run("Delete Product", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/product/"+ProductId, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("Get Product After Delete", func(t *testing.T) {
		resp := GlobalGetProductId(t, app, ProductId)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		var fetchedProduct FetchedResponse[domain.Product]
		json.NewDecoder(resp.Body).Decode(&fetchedProduct)
		assert.Equal(t, http.StatusInternalServerError, fetchedProduct.Code)
		assert.Equal(t, "product not found", fetchedProduct.Message)
		assert.Empty(t, fetchedProduct.Data)
	})
}
