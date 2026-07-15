package usecase

import (
	"simple_api/internal/domain/entity"
	"simple_api/internal/domain/repository"
)

type BookUseCase interface {
	GetById(id int) (entity.Book, error)
	Create(book entity.Book) (entity.Book, error)
	Update(id int, book entity.Book) (entity.Book, error)
	Delete(id int) error
}

type book struct {
	repo repository.BookRepository
}

func NewBook(repo repository.BookRepository) BookUseCase {
	return &book{repo: repo}
}

func (uc *book) GetById(id int) (entity.Book, error) {
	return uc.repo.GetById(id)
}

func (uc *book) Create(book entity.Book) (entity.Book, error) {
	return uc.repo.Create(book)
}

func (uc *book) Update(id int, book entity.Book) (entity.Book, error) {
	book.Id = id
	return uc.repo.Update(book)
}

func (uc *book) Delete(id int) error {
	return uc.repo.Delete(id)
}
