package database

import (
	"simple_api/app/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New() (db *gorm.DB, err error) {
	dsn := config.Cfg.Dsn()

	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, err
}
