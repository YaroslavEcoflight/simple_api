package database

import (
	"simple_api/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(config.Cfg.Dsn()), &gorm.Config{})
}
