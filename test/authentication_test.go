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
	"github.com/stretchr/testify/assert"
)

func deleteTestUser(conn *sql.DB) {
	conn.Exec("DELETE FROM users WHERE name='test'")
}

func AuthLoginSetup() *fiber.App {
	app := fiber.New()
	conn := GlobalSetupTest()
	deleteTestUser(conn)
	CreateAdminTest(conn)
	conf := InitializedLoader()

	api.NewAuthApi(app, conn, &conf.Jwt)

	return app
}

func TestLoginSuccess(t *testing.T) {
	app := AuthLoginSetup()
	loginData := dto.SignInRequest{
		Username: "admin@admin.com",
		Password: "12345678",
	}

	jsonData, _ := json.Marshal(loginData)

	req := httptest.NewRequest(http.MethodPost, "/auth", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, _ := app.Test(req)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Login failed, expected %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var fetchedAuth FetchedResponse[dto.SignInResponse]
	json.NewDecoder(resp.Body).Decode(&fetchedAuth)

	assert.Equal(t, "success", fetchedAuth.Message)

	if fetchedAuth.Data.Token == "" {
		t.Errorf("Login failed, expected %s, got %s", "token", fetchedAuth.Data.Token)
	}

	assert.NotEmpty(t, fetchedAuth.Data.Token)

}

func TestLoginFail(t *testing.T) {
	app := AuthLoginSetup()

	t.Run("LoginWithWrongPassword", func(t *testing.T) {
		loginData := dto.SignInRequest{
			Username: "admin@admin.com",
			Password: "12345671",
		}

		jsonData, _ := json.Marshal(loginData)

		req := httptest.NewRequest(http.MethodPost, "/auth", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("Login failed, expected %d, got %d", http.StatusOK, resp.StatusCode)
		}

		var fetchedAuth FetchedResponse[dto.SignInResponse]
		json.NewDecoder(resp.Body).Decode(&fetchedAuth)

		assert.Equal(t, 401, resp.StatusCode)

		assert.Equal(t, "username or password incorrect", fetchedAuth.Message)

		assert.Empty(t, fetchedAuth.Data.Token)
	})

	t.Run("LoginWithWrongEmail", func(t *testing.T) {
		loginData := dto.SignInRequest{
			Username: "admins@admin.com",
			Password: "12345678",
		}

		jsonData, _ := json.Marshal(loginData)

		req := httptest.NewRequest(http.MethodPost, "/auth", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("Login failed, expected %d, got %d", http.StatusOK, resp.StatusCode)
		}

		var fetchedAuth FetchedResponse[dto.SignInResponse]
		json.NewDecoder(resp.Body).Decode(&fetchedAuth)

		assert.Equal(t, 401, resp.StatusCode)

		assert.Equal(t, "username or password incorrect", fetchedAuth.Message)

		assert.Empty(t, fetchedAuth.Data.Token)
	})

	t.Run("LoginWithWrongEmailAndPassword", func(t *testing.T) {
		loginData := dto.SignInRequest{
			Username: "admins@admin.com",
			Password: "1234567q",
		}

		jsonData, _ := json.Marshal(loginData)

		req := httptest.NewRequest(http.MethodPost, "/auth", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		resp, _ := app.Test(req)

		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("Login failed, expected %d, got %d", http.StatusOK, resp.StatusCode)
		}

		var fetchedAuth FetchedResponse[dto.SignInResponse]
		json.NewDecoder(resp.Body).Decode(&fetchedAuth)

		assert.Equal(t, 401, resp.StatusCode)

		assert.Equal(t, "username or password incorrect", fetchedAuth.Message)

		assert.Empty(t, fetchedAuth.Data.Token)
	})

}
