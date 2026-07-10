package model

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	Id        int    `gorm:"column:id;primaryKey"`
	Title     string `gorm:"column:title"`
	Author    string `gorm:"column:author"`
	Rating    int    `gorm:"column:rating"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Book) TableName() string {
	return "books"
}

func (m *Book) BeforeCreate(tx *gorm.DB) {
	// Видел как uuid заполняют, нужно ли такое делать для int
	// ответ - не нужно
}
