package usecase

import (
	"simple_api/internal/domain/entity"
	"simple_api/internal/domain/repository"
)

type BookUseCase interface {
	GetById(id int) (entity.Book, error)
	Create(title, author string, rating int) (entity.Book, error)
	Update(id int, title, author string, rating int) (entity.Book, error)
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

func (uc *book) Create(title, author string, rating int) (entity.Book, error) {
	return uc.repo.Create(entity.Book{Title: title, Author: author, Rating: rating})
}

func (uc *book) Update(id int, title, author string, rating int) (entity.Book, error) {
	return uc.repo.Update(entity.Book{Id: id, Title: title, Author: author, Rating: rating})
}

func (uc *book) Delete(id int) error {
	return uc.repo.Delete(id)
}
