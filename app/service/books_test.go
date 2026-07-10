package service

import (
	"testing"

	"simple_api/app/dto"
	"simple_api/app/model"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type mockBookRepo struct {
	getBookByIdFn func(id int) (model.Book, error)
	createBookFn  func(entity model.Book) (model.Book, error)
	updateBookFn  func(entity model.Book) (model.Book, error)
	deleteBookFn  func(id int) error
}

func (m *mockBookRepo) GetBookById(id int) (model.Book, error) {
	if m.getBookByIdFn != nil {
		return m.getBookByIdFn(id)
	}
	return model.Book{Id: id, Title: "test"}, nil
}

func (m *mockBookRepo) CreateBook(entity model.Book) (model.Book, error) {
	if m.createBookFn != nil {
		return m.createBookFn(entity)
	}
	return entity, nil
}

func (m *mockBookRepo) UpdateBook(entity model.Book) (model.Book, error) {
	if m.updateBookFn != nil {
		return m.updateBookFn(entity)
	}
	return entity, nil
}

func (m *mockBookRepo) DeleteBook(id int) error {
	if m.deleteBookFn != nil {
		return m.deleteBookFn(id)
	}
	return nil
}

func TestGetBookById(t *testing.T) {
	svc := &book{repo: &mockBookRepo{}}

	result, err := svc.GetBookById(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, result.Id)
	assert.Equal(t, "test", result.Title)
}

func TestGetBookById_NotFound(t *testing.T) {
	svc := &book{repo: &mockBookRepo{
		getBookByIdFn: func(id int) (model.Book, error) {
			return model.Book{}, gorm.ErrRecordNotFound
		},
	}}
	_, err := svc.GetBookById(999)
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
}

func TestCreateBook(t *testing.T) {
	svc := &book{repo: &mockBookRepo{}}
	entity := dto.BookCreateRequest{
		Title:  "test",
		Author: "test_author",
		Rating: 1,
	}
	result, err := svc.CreateBook(entity)
	assert.NoError(t, err)
	assert.Equal(t, "test", result.Title)
	assert.Equal(t, "test_author", result.Author)
}

func TestUpdateBook(t *testing.T) {
	svc := &book{repo: &mockBookRepo{}}
	entity := dto.BookUpdateRequest{
		Id:     1,
		Title:  "test",
		Author: "test_author",
		Rating: 1,
	}
	result, err := svc.UpdateBook(entity)

	assert.NoError(t, err)
	assert.Equal(t, "test", result.Title)
	assert.Equal(t, "test_author", result.Author)
	assert.Equal(t, 1, result.Rating)
}

func TestDeleteBook(t *testing.T) {
	svc := &book{repo: &mockBookRepo{}}

	id := 1

	err := svc.DeleteBook(id)
	assert.NoError(t, err)
}

func TestDeleteBook_NotFound(t *testing.T) {
	svc := &book{
		repo: &mockBookRepo{
			deleteBookFn: func(id int) error {
				return gorm.ErrRecordNotFound
			},
		},
	}
	err := svc.DeleteBook(999)
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
}
