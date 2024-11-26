package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/pewe21/PointOfSale/api"
	"github.com/pewe21/PointOfSale/dto"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const CUSTOMER_URL = "/customer"

func deleteAllCustomer(conn *sql.DB) {
	conn.Exec("DELETE FROM customers")
}

func CustomerSetup() *fiber.App {
	app := fiber.New()
	conn := GlobalSetupTest()
	deleteAllCustomer(conn)
	api.NewCustomerApi(app, conn)

	return app
}

func GlobalCreateCustomer(t *testing.T, app *fiber.App, role dto.CreateCustomerRequest) *http.Response {
	body, _ := json.Marshal(role)

	req := httptest.NewRequest(http.MethodPost, CUSTOMER_URL, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	return resp
}

func GlobalGetCustomer(t *testing.T, app *fiber.App) *http.Response {
	req := httptest.NewRequest(http.MethodGet, CUSTOMER_URL, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, _ := app.Test(req)

	return resp
}

func GlobalGetCustomerId(t *testing.T, app *fiber.App, id string) *http.Response {
	req := httptest.NewRequest(http.MethodGet, CUSTOMER_URL+"/"+id, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, _ := app.Test(req)

	return resp
}

func TestCreateCustomer(t *testing.T) {
	app := CustomerSetup()

	t.Run("CreateCustomerSuccess", func(t *testing.T) {
		dataCustomer := dto.CreateCustomerRequest{
			Name:  "test",
			Phone: "081123123123",
		}

		resp := GlobalCreateCustomer(t, app, dataCustomer)

		var fetchedResponseCustomer FetchedResponse[dto.CustomerData]

		json.NewDecoder(resp.Body).Decode(&fetchedResponseCustomer)

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
		}

		assert.Equal(t, "created successfully", fetchedResponseCustomer.Message)
		assert.Empty(t, fetchedResponseCustomer.Data)
	})

	t.Run("CreateCustomerFailed", func(t *testing.T) {
		t.Run("CreateCustomerBlankName", func(t *testing.T) {
			dataCustomer := dto.CreateCustomerRequest{
				Name:  "",
				Phone: "081123123123",
			}

			resp := GlobalCreateCustomer(t, app, dataCustomer)

			var fetchedResponseCustomer FetchedResponse[dto.CustomerData]

			json.NewDecoder(resp.Body).Decode(&fetchedResponseCustomer)

			if resp.StatusCode != http.StatusBadRequest {
				t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
			}
			assert.Empty(t, fetchedResponseCustomer.Data)
		})
		t.Run("CreateCustomerBlankName", func(t *testing.T) {
			dataCustomer := dto.CreateCustomerRequest{
				Name:  "",
				Phone: "081123123123",
			}

			resp := GlobalCreateCustomer(t, app, dataCustomer)

			var fetchedResponseCustomer FetchedResponse[dto.CustomerData]

			json.NewDecoder(resp.Body).Decode(&fetchedResponseCustomer)

			if resp.StatusCode != http.StatusBadRequest {
				t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
			}

			assert.Empty(t, fetchedResponseCustomer.Data)
		})
	})
}

func TestGetAllCustomer(t *testing.T) {
	app := CustomerSetup()

	dataCustomer := dto.CreateCustomerRequest{
		Name:  "test",
		Phone: "081123123123",
	}

	respCreate := GlobalCreateCustomer(t, app, dataCustomer)

	if respCreate.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, respCreate.StatusCode)
	}

	t.Run("GetAllCustomerSuccess", func(t *testing.T) {
		customer := GlobalGetCustomer(t, app)
		var fetcherCustomer FetchedResponse[[]dto.CustomerData]

		if customer.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, customer.StatusCode)
		}

		json.NewDecoder(customer.Body).Decode(&fetcherCustomer)

		assert.NotEmpty(t, fetcherCustomer.Data)
		assert.Equal(t, "success", fetcherCustomer.Message)

	})
}

func TestUpdateCustomer(t *testing.T) {
	app := CustomerSetup()
	dataCustomer := dto.CreateCustomerRequest{
		Name:  "test",
		Phone: "081123123123",
	}

	respCreate := GlobalCreateCustomer(t, app, dataCustomer)

	if respCreate.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, respCreate.StatusCode)
	}

	getCustomer := GlobalGetCustomer(t, app)

	if getCustomer.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, getCustomer.StatusCode)
	}

	var fetchedResponseCustomer FetchedResponse[[]dto.CustomerData]

	json.NewDecoder(getCustomer.Body).Decode(&fetchedResponseCustomer)

	assert.NotEmpty(t, fetchedResponseCustomer.Data)

	var fetchedResponseCustomerUpdate FetchedResponse[dto.CustomerData]
	t.Run("UpdateCustomerSuccess", func(t *testing.T) {

		dataUpdateCustomer := dto.UpdateCustomerRequest{
			Name:    "testing",
			Phone:   "081123123123",
			Address: "cekcek",
		}

		body, _ := json.Marshal(dataUpdateCustomer)

		req := httptest.NewRequest(http.MethodPut, CUSTOMER_URL+"/"+fetchedResponseCustomer.Data[0].Id, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		respUpdate, _ := app.Test(req)

		json.NewDecoder(respUpdate.Body).Decode(&fetchedResponseCustomerUpdate)

		if respUpdate.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, respUpdate.StatusCode)
		}

		assert.Equal(t, "success", fetchedResponseCustomerUpdate.Message)
		assert.Empty(t, fetchedResponseCustomerUpdate.Data)

		t.Run("check update data customer", func(t *testing.T) {
			customerData := GlobalGetCustomerId(t, app, fetchedResponseCustomer.Data[0].Id)

			json.NewDecoder(customerData.Body).Decode(&fetchedResponseCustomerUpdate)
			if customerData.StatusCode != http.StatusOK {
				t.Errorf("Expected status code %d, got %d", http.StatusOK, customerData.StatusCode)
			}

			assert.Equal(t, "success", fetchedResponseCustomerUpdate.Message)
			assert.NotEqual(t, "", fetchedResponseCustomerUpdate.Data.Id)
			assert.Equal(t, "testing", fetchedResponseCustomerUpdate.Data.Name)
			assert.Equal(t, "081123123123", fetchedResponseCustomerUpdate.Data.Phone)
		})
	})

	t.Run("UpdateCustomerFailed", func(t *testing.T) {

		dataUpdateCustomer := dto.UpdateCustomerRequest{
			Name:    "",
			Phone:   "081123123123",
			Address: "cekcek",
		}

		body, _ := json.Marshal(dataUpdateCustomer)

		req := httptest.NewRequest(http.MethodPut, CUSTOMER_URL+"/"+fetchedResponseCustomer.Data[0].Id, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		respUpdate, _ := app.Test(req)

		var fetchedResponseCustomerUpdateErr FetchedResponse[dto.CustomerData]

		json.NewDecoder(respUpdate.Body).Decode(&fetchedResponseCustomerUpdateErr)

		if respUpdate.StatusCode != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, respUpdate.StatusCode)
		}

		assert.Contains(t, fetchedResponseCustomerUpdateErr.Message, "required")
		assert.Empty(t, fetchedResponseCustomerUpdateErr.Data)
	})
}

func TestDeleteCustomer(t *testing.T) {
	app := CustomerSetup()

	t.Run("DeleteCustomerSuccess", func(t *testing.T) {
		dataCustomer := dto.CreateCustomerRequest{
			Name:  "test",
			Phone: "081123123123",
		}
		customer := GlobalCreateCustomer(t, app, dataCustomer)

		if customer.StatusCode != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, customer.StatusCode)
		}
		getCustomer := GlobalGetCustomer(t, app)

		if getCustomer.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, getCustomer.StatusCode)
		}

		var fetchedCustomersData FetchedResponse[[]dto.CustomerData]

		json.NewDecoder(getCustomer.Body).Decode(&fetchedCustomersData)

		assert.NotEmpty(t, fetchedCustomersData.Data)

		req := httptest.NewRequest(http.MethodDelete, CUSTOMER_URL+"/"+fetchedCustomersData.Data[0].Id, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		respDelete, _ := app.Test(req)

		if respDelete.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, respDelete.StatusCode)
		}

		var fetchedCustomerResonseDelete FetchedResponse[dto.CustomerData]
		json.NewDecoder(respDelete.Body).Decode(&fetchedCustomerResonseDelete)

		assert.Equal(t, "success", fetchedCustomerResonseDelete.Message)
		assert.Empty(t, fetchedCustomerResonseDelete.Data)

		t.Run("check delete data customer", func(t *testing.T) {
			customerData := GlobalGetCustomerId(t, app, fetchedCustomersData.Data[0].Id)
			if customerData.StatusCode != http.StatusOK {
				t.Errorf("Expected status code %d, got %d", http.StatusOK, customerData.StatusCode)
			}
			var fetchedCheckDataAfterDelete FetchedResponse[dto.CustomerData]
			json.NewDecoder(customerData.Body).Decode(&fetchedCheckDataAfterDelete)
			assert.Equal(t, "success", fetchedCheckDataAfterDelete.Message)
			assert.Empty(t, fetchedCheckDataAfterDelete.Data)
		})

	})

	t.Run("DeleteCustomerFailed", func(t *testing.T) {
		dataCustomer := dto.CreateCustomerRequest{
			Name:  "test",
			Phone: "081123123123",
		}
		customer := GlobalCreateCustomer(t, app, dataCustomer)

		if customer.StatusCode != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, customer.StatusCode)
		}

		getCustomer := GlobalGetCustomer(t, app)

		if getCustomer.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, getCustomer.StatusCode)
		}

		var fetchedCustomersData FetchedResponse[[]dto.CustomerData]

		json.NewDecoder(getCustomer.Body).Decode(&fetchedCustomersData)

		assert.NotEmpty(t, fetchedCustomersData.Data)

		req := httptest.NewRequest(http.MethodDelete, CUSTOMER_URL+"/0", nil)

		respDelete, _ := app.Test(req)

		if respDelete.StatusCode != http.StatusNotFound {
			t.Errorf("Expected status code %d, got %d", http.StatusNotFound, respDelete.StatusCode)
		}

		var fetchedCustomerResponseDelete FetchedResponse[dto.CustomerData]
		json.NewDecoder(respDelete.Body).Decode(&fetchedCustomerResponseDelete)
		assert.Equal(t, "customer not found", fetchedCustomerResponseDelete.Message)
		assert.Empty(t, fetchedCustomerResponseDelete.Data)

		t.Run("check delete data customer", func(t *testing.T) {
			customerData := GlobalGetCustomerId(t, app, fetchedCustomersData.Data[0].Id)
			if customerData.StatusCode != http.StatusOK {
				t.Errorf("Expected status code %d, got %d", http.StatusOK, customerData.StatusCode)
			}
			var fetchedCustomerCheckDataEmpty FetchedResponse[dto.CustomerData]
			json.NewDecoder(customerData.Body).Decode(&fetchedCustomerCheckDataEmpty)
			assert.Equal(t, "success", fetchedCustomerCheckDataEmpty.Message)
			assert.NotEmpty(t, fetchedCustomerCheckDataEmpty.Data)
		})

	})
}
