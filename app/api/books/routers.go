package books

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AddRouters(r fiber.Router, db *gorm.DB) {
	handler := NewHandler(db)
	router := r.Group("/book")
	router.Get("/:id", handler.GetById)
	router.Post("", handler.Create)
	router.Put("/:id", handler.Update)
	router.Delete("/:id", handler.Delete)
}
