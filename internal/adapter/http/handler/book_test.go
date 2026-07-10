package handler

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"simple_api/internal/domain"
	"simple_api/internal/domain/entity"
	"simple_api/internal/usecase"
)

type mockBookUseCase struct {
	getByIdFn func(id int) (entity.Book, error)
	createFn  func(title, author string, rating int) (entity.Book, error)
	updateFn  func(id int, title, author string, rating int) (entity.Book, error)
	deleteFn  func(id int) error
}

func (m *mockBookUseCase) GetById(id int) (entity.Book, error) {
	if m.getByIdFn != nil {
		return m.getByIdFn(id)
	}
	return entity.Book{Id: id, Title: "test"}, nil
}

func (m *mockBookUseCase) Create(title, author string, rating int) (entity.Book, error) {
	if m.createFn != nil {
		return m.createFn(title, author, rating)
	}
	return entity.Book{Title: title, Author: author, Rating: rating}, nil
}

func (m *mockBookUseCase) Update(id int, title, author string, rating int) (entity.Book, error) {
	if m.updateFn != nil {
		return m.updateFn(id, title, author, rating)
	}
	return entity.Book{Id: id, Title: title, Author: author, Rating: rating}, nil
}

func (m *mockBookUseCase) Delete(id int) error {
	if m.deleteFn != nil {
		return m.deleteFn(id)
	}
	return nil
}

func newTestApp(uc usecase.BookUseCase) *fiber.App {
	app := fiber.New()
	h := NewBookHandler(uc)
	app.Get("/book/:id", h.GetById)
	app.Post("/book", h.Create)
	app.Put("/book/:id", h.Update)
	app.Delete("/book/:id", h.Delete)
	return app
}

func TestGetById(t *testing.T) {
	app := newTestApp(&mockBookUseCase{})

	req := httptest.NewRequest("GET", "/book/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetById_NotFound(t *testing.T) {
	app := newTestApp(&mockBookUseCase{
		getByIdFn: func(id int) (entity.Book, error) {
			return entity.Book{}, domain.ErrNotFound
		},
	})

	req := httptest.NewRequest("GET", "/book/999", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 404, resp.StatusCode)
}

func TestCreate(t *testing.T) {
	app := newTestApp(&mockBookUseCase{})

	body := strings.NewReader(`{"title":"test","author":"author","rating":5}`)
	req := httptest.NewRequest("POST", "/book", body)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, 201, resp.StatusCode)
}

func TestCreate_InvalidBody(t *testing.T) {
	app := newTestApp(&mockBookUseCase{})

	req := httptest.NewRequest("POST", "/book", strings.NewReader(`invalid`))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, 400, resp.StatusCode)
}

func TestUpdate(t *testing.T) {
	app := newTestApp(&mockBookUseCase{})

	body := strings.NewReader(`{"title":"updated","author":"author","rating":5}`)
	req := httptest.NewRequest("PUT", "/book/1", body)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestUpdate_NotFound(t *testing.T) {
	app := newTestApp(&mockBookUseCase{
		updateFn: func(id int, title, author string, rating int) (entity.Book, error) {
			return entity.Book{}, domain.ErrNotFound
		},
	})

	body := strings.NewReader(`{"title":"updated","author":"author","rating":5}`)
	req := httptest.NewRequest("PUT", "/book/999", body)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, 404, resp.StatusCode)
}

func TestDelete(t *testing.T) {
	app := newTestApp(&mockBookUseCase{})

	req := httptest.NewRequest("DELETE", "/book/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestDelete_NotFound(t *testing.T) {
	app := newTestApp(&mockBookUseCase{
		deleteFn: func(id int) error {
			return domain.ErrNotFound
		},
	})

	req := httptest.NewRequest("DELETE", "/book/999", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 404, resp.StatusCode)
}
