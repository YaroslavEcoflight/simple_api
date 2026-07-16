package e2e

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	httpadapter "simple_api/internal/adapter/http"
	adapterrepo "simple_api/internal/adapter/repository"
	"simple_api/internal/infrastructure/model"
	"simple_api/internal/usecase"
)

func setupApp(t *testing.T) *fiber.App {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&model.Book{}))
	t.Cleanup(func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	})

	repo := adapterrepo.NewBookRepository(db)
	uc := usecase.NewBook(repo)

	app := fiber.New()
	api := app.Group("/api/v1")
	httpadapter.RegisterRoutes(api, uc)
	return app
}

func TestCreateBook(t *testing.T) {
	app := setupApp(t)

	body := strings.NewReader(`{"title":"Go Programming","author":"Donovan","rating":5}`)
	req := httptest.NewRequest("POST", "/api/v1/book", body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	require.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)
}

func TestCreateBook_InvalidBody(t *testing.T) {
	app := setupApp(t)

	req := httptest.NewRequest("POST", "/api/v1/book", strings.NewReader(`invalid json`))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	require.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestGetBook(t *testing.T) {
	app := setupApp(t)

	body := strings.NewReader(`{"title":"Go Programming","author":"Donovan","rating":5}`)
	createReq := httptest.NewRequest("POST", "/api/v1/book", body)
	createReq.Header.Set("Content-Type", "application/json")
	_, err := app.Test(createReq)
	require.NoError(t, err)

	req := httptest.NewRequest("GET", "/api/v1/book/1", nil)
	resp, err := app.Test(req)

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetBook_NotFound(t *testing.T) {
	app := setupApp(t)

	req := httptest.NewRequest("GET", "/api/v1/book/999", nil)
	resp, err := app.Test(req)

	require.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestGetBook_InvalidId(t *testing.T) {
	app := setupApp(t)

	req := httptest.NewRequest("GET", "/api/v1/book/abc", nil)
	resp, err := app.Test(req)

	require.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestUpdateBook(t *testing.T) {
	app := setupApp(t)

	createBody := strings.NewReader(`{"title":"Go Programming","author":"Donovan","rating":5}`)
	createReq := httptest.NewRequest("POST", "/api/v1/book", createBody)
	createReq.Header.Set("Content-Type", "application/json")
	_, err := app.Test(createReq)
	require.NoError(t, err)

	updateBody := strings.NewReader(`{"title":"Go Programming 2nd Ed","author":"Donovan","rating":5}`)
	req := httptest.NewRequest("PUT", "/api/v1/book/1", updateBody)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUpdateBook_NotFound(t *testing.T) {
	app := setupApp(t)

	body := strings.NewReader(`{"title":"Ghost","author":"Nobody","rating":1}`)
	req := httptest.NewRequest("PUT", "/api/v1/book/999", body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	require.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUpdateBook_InvalidId(t *testing.T) {
	app := setupApp(t)

	body := strings.NewReader(`{"title":"Ghost","author":"Nobody","rating":1}`)
	req := httptest.NewRequest("PUT", "/api/v1/book/abc", body)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	require.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestDeleteBook(t *testing.T) {
	app := setupApp(t)

	createBody := strings.NewReader(`{"title":"Go Programming","author":"Donovan","rating":5}`)
	createReq := httptest.NewRequest("POST", "/api/v1/book", createBody)
	createReq.Header.Set("Content-Type", "application/json")
	_, err := app.Test(createReq)
	require.NoError(t, err)

	req := httptest.NewRequest("DELETE", "/api/v1/book/1", nil)
	resp, err := app.Test(req)

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestDeleteBook_NotFound(t *testing.T) {
	app := setupApp(t)

	req := httptest.NewRequest("DELETE", "/api/v1/book/999", nil)
	resp, err := app.Test(req)

	require.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestDeleteBook_InvalidId(t *testing.T) {
	app := setupApp(t)

	req := httptest.NewRequest("DELETE", "/api/v1/book/abc", nil)
	resp, err := app.Test(req)

	require.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)
}
