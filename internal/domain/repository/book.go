package repository

import "simple_api/internal/domain/entity"

type BookRepository interface {
	GetById(id int) (entity.Book, error)
	Create(book entity.Book) (entity.Book, error)
	Update(book entity.Book) (entity.Book, error)
	Delete(id int) error
}
