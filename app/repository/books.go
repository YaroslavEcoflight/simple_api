package repository

import (
	"gorm.io/gorm"

	"simple_api/app/model"
)

type Book interface {
	GetBookById(id int) (model.Book, error)
	CreateBook(entity model.Book) (model.Book, error)
	UpdateBook(etity model.Book) (model.Book, error)
	DeleteBook(id int) error
}

type book struct {
	db *gorm.DB
}

func NewBook(db *gorm.DB) *book {
	return &book{db: db}
}

func (b *book) GetBookById(id int) (entity model.Book, err error) {
	return entity, b.db.First(&entity, "id = ?", id).Error
}

func (b *book) CreateBook(entity model.Book) (model.Book, error) {
	// Создание книги
	return entity, b.db.Create(&entity).Error
}

func (b *book) UpdateBook(entity model.Book) (model.Book, error) {
	existing, err := b.GetBookById(entity.Id)
	if err != nil {
		return existing, err
	}
	return existing, b.db.Model(&existing).Updates(entity).Error
}

func (b *book) DeleteBook(id int) error {
	result := b.db.Delete(&model.Book{Id: id})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
