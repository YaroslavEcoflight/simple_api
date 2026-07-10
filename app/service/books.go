package service

import (
	"gorm.io/gorm"

	"simple_api/app/dto"
	"simple_api/app/model"
	"simple_api/app/repository"
)

type Book interface {
	GetBookById(id int) (model.Book, error)
	CreateBook(entity dto.BookCreateRequest) (model.Book, error)
	UpdateBook(entity dto.BookUpdateRequest) (model.Book, error)
	DeleteBook(id int) error
}

type book struct {
	repo repository.Book
}

func NewBook(db *gorm.DB) *book {
	repo := repository.NewBook(db)
	return &book{repo: repo}
}

func (b *book) GetBookById(id int) (model.Book, error) {
	return b.repo.GetBookById(id)
}

func (b *book) CreateBook(entity dto.BookCreateRequest) (model.Book, error) {
	return b.repo.CreateBook(
		model.Book{
			Title:  entity.Title,
			Author: entity.Author,
			Rating: entity.Rating,
		},
	)
}

func (b *book) UpdateBook(entity dto.BookUpdateRequest) (model.Book, error) {
	return b.repo.UpdateBook(
		model.Book{
			Id:     entity.Id,
			Title:  entity.Title,
			Author: entity.Author,
			Rating: entity.Rating,
		},
	)

}

func (b *book) DeleteBook(id int) error {
	return b.repo.DeleteBook(id)
}
