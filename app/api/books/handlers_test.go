package books

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"simple_api/app/dto"
	"simple_api/app/model"
	"simple_api/app/service"
)

type mockBookSvc struct {
	getBookByIdFn func(id int) (model.Book, error)
	createBookFn  func(entity dto.BookCreateRequest) (model.Book, error)
	updateBookFn  func(entity dto.BookUpdateRequest) (model.Book, error)
	deleteBookFn  func(id int) error
}

func (m *mockBookSvc) GetBookById(id int) (model.Book, error) {
	if m.getBookByIdFn != nil {
		return m.getBookByIdFn(id)
	}
	return model.Book{Id: id, Title: "test"}, nil
}

func (m *mockBookSvc) CreateBook(entity dto.BookCreateRequest) (model.Book, error) {
	if m.createBookFn != nil {
		return m.createBookFn(entity)
	}
	return model.Book{Title: entity.Title, Author: entity.Author, Rating: entity.Rating}, nil
}

func (m *mockBookSvc) UpdateBook(entity dto.BookUpdateRequest) (model.Book, error) {
	if m.updateBookFn != nil {
		return m.updateBookFn(entity)
	}
	return model.Book{Id: entity.Id, Title: entity.Title, Author: entity.Author, Rating: entity.Rating}, nil
}

func (m *mockBookSvc) DeleteBook(id int) error {
	if m.deleteBookFn != nil {
		return m.deleteBookFn(id)
	}
	return nil
}

func newTestApp(svc service.Book) *fiber.App {
	app := fiber.New()
	h := &Handler{svc: svc}
	app.Get("/book/:id", h.GetById)
	app.Post("/book", h.Create)
	app.Put("/book/:id", h.Update)
	app.Delete("/book/:id", h.Delete)
	return app
}

func TestGetBookById(t *testing.T) {
	app := newTestApp(&mockBookSvc{})

	req := httptest.NewRequest("GET", "/book/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetBookById_NotFound(t *testing.T) {
	app := newTestApp(&mockBookSvc{
		getBookByIdFn: func(id int) (model.Book, error) {
			return model.Book{}, gorm.ErrRecordNotFound
		},
	})

	req := httptest.NewRequest("GET", "/book/999", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 404, resp.StatusCode)
}

func TestCreate(t *testing.T) {
	app := newTestApp(&mockBookSvc{})

	body := strings.NewReader(`{"title":"test","author":"author","rating":5}`)
	req := httptest.NewRequest("POST", "/book", body)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, 201, resp.StatusCode)
}

func TestCreate_InvalidBody(t *testing.T) {
	app := newTestApp(&mockBookSvc{})

	req := httptest.NewRequest("POST", "/book", strings.NewReader(`invalid`))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, 400, resp.StatusCode)
}

func TestUpdate(t *testing.T) {
	app := newTestApp(&mockBookSvc{})

	body := strings.NewReader(`{"title":"updated","author":"author","rating":5}`)
	req := httptest.NewRequest("PUT", "/book/1", body)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestUpdate_NotFound(t *testing.T) {
	app := newTestApp(&mockBookSvc{
		updateBookFn: func(entity dto.BookUpdateRequest) (model.Book, error) {
			return model.Book{}, gorm.ErrRecordNotFound
		},
	})

	body := strings.NewReader(`{"title":"updated","author":"author","rating":5}`)
	req := httptest.NewRequest("PUT", "/book/999", body)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, 404, resp.StatusCode)
}

func TestDelete(t *testing.T) {
	app := newTestApp(&mockBookSvc{})

	req := httptest.NewRequest("DELETE", "/book/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestDelete_NotFound(t *testing.T) {
	app := newTestApp(&mockBookSvc{
		deleteBookFn: func(id int) error {
			return gorm.ErrRecordNotFound
		},
	})

	req := httptest.NewRequest("DELETE", "/book/999", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 404, resp.StatusCode)
}
