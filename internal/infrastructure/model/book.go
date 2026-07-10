package model

import "time"

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
