package usecase

import (
	"testing"

	"simple_api/internal/domain"
	"simple_api/internal/domain/entity"

	"github.com/stretchr/testify/assert"
)

type mockBookRepo struct {
	getByIdFn func(id int) (entity.Book, error)
	createFn  func(book entity.Book) (entity.Book, error)
	updateFn  func(book entity.Book) (entity.Book, error)
	deleteFn  func(id int) error
}

func (m *mockBookRepo) GetById(id int) (entity.Book, error) {
	if m.getByIdFn != nil {
		return m.getByIdFn(id)
	}
	return entity.Book{Id: id, Title: "test"}, nil
}

func (m *mockBookRepo) Create(b entity.Book) (entity.Book, error) {
	if m.createFn != nil {
		return m.createFn(b)
	}
	return b, nil
}

func (m *mockBookRepo) Update(b entity.Book) (entity.Book, error) {
	if m.updateFn != nil {
		return m.updateFn(b)
	}
	return b, nil
}

func (m *mockBookRepo) Delete(id int) error {
	if m.deleteFn != nil {
		return m.deleteFn(id)
	}
	return nil
}

func TestGetById(t *testing.T) {
	uc := NewBook(&mockBookRepo{})
	result, err := uc.GetById(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, result.Id)
	assert.Equal(t, "test", result.Title)
}

func TestGetById_NotFound(t *testing.T) {
	uc := NewBook(&mockBookRepo{
		getByIdFn: func(id int) (entity.Book, error) {
			return entity.Book{}, domain.ErrNotFound
		},
	})
	_, err := uc.GetById(999)
	assert.ErrorIs(t, err, domain.ErrNotFound)
}

func TestCreate(t *testing.T) {
	uc := NewBook(&mockBookRepo{})
	result, err := uc.Create(entity.Book{Title: "test", Author: "test_author", Rating: 1})
	assert.NoError(t, err)
	assert.Equal(t, "test", result.Title)
	assert.Equal(t, "test_author", result.Author)
}

func TestUpdate(t *testing.T) {
	uc := NewBook(&mockBookRepo{})
	result, err := uc.Update(1, entity.Book{Title: "test", Author: "test_author", Rating: 1})
	assert.NoError(t, err)
	assert.Equal(t, "test", result.Title)
	assert.Equal(t, "test_author", result.Author)
	assert.Equal(t, 1, result.Rating)
}

func TestDelete(t *testing.T) {
	uc := NewBook(&mockBookRepo{})
	err := uc.Delete(1)
	assert.NoError(t, err)
}

func TestDelete_NotFound(t *testing.T) {
	uc := NewBook(&mockBookRepo{
		deleteFn: func(id int) error {
			return domain.ErrNotFound
		},
	})
	err := uc.Delete(999)
	assert.ErrorIs(t, err, domain.ErrNotFound)
}
