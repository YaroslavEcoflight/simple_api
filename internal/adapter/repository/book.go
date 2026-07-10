package repository

import (
	"errors"

	"gorm.io/gorm"

	"simple_api/internal/domain"
	"simple_api/internal/domain/entity"
	domainrepo "simple_api/internal/domain/repository"
	"simple_api/internal/infrastructure/model"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) domainrepo.BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) GetById(id int) (entity.Book, error) {
	var m model.Book
	if err := r.db.First(&m, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Book{}, domain.ErrNotFound
		}
		return entity.Book{}, err
	}
	return toEntity(m), nil
}

func (r *bookRepository) Create(book entity.Book) (entity.Book, error) {
	m := toModel(book)
	if err := r.db.Create(&m).Error; err != nil {
		return entity.Book{}, err
	}
	return toEntity(m), nil
}

func (r *bookRepository) Update(book entity.Book) (entity.Book, error) {
	var existing model.Book
	if err := r.db.First(&existing, "id = ?", book.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Book{}, domain.ErrNotFound
		}
		return entity.Book{}, err
	}
	if err := r.db.Model(&existing).Updates(toModel(book)).Error; err != nil {
		return entity.Book{}, err
	}
	return toEntity(existing), nil
}

func (r *bookRepository) Delete(id int) error {
	result := r.db.Delete(&model.Book{Id: id})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func toEntity(m model.Book) entity.Book {
	return entity.Book{Id: m.Id, Title: m.Title, Author: m.Author, Rating: m.Rating}
}

func toModel(e entity.Book) model.Book {
	return model.Book{Id: e.Id, Title: e.Title, Author: e.Author, Rating: e.Rating}
}
