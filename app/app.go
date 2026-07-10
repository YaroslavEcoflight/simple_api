package app

import (
	"simple_api/config"
	httpadapter "simple_api/internal/adapter/http"
	adapterrepo "simple_api/internal/adapter/repository"
	"simple_api/internal/infrastructure/database"
	"simple_api/internal/infrastructure/model"
	"simple_api/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	db, err := database.New()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Book{})

	repo := adapterrepo.NewBookRepository(db)
	uc := usecase.NewBook(repo)

	app := fiber.New()
	api := app.Group("/api/v1")
	httpadapter.RegisterRoutes(api, uc)

	app.Listen(config.Cfg.Addr())
}
