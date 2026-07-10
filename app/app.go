package app

import (
	"simple_api/app/api/books"
	"simple_api/app/model"
	"simple_api/pkg/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Run() {
	db, err := database.New()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Book{})
	app := fiber.New(fiber.Config{})
	api := app.Group("/api")

	initAPIV1(api, db)

	app.Listen(":3000")

}

func initAPIV1(app fiber.Router, db *gorm.DB) {
	v1 := app.Group("v1")
	books.AddRouters(v1, db)
}
