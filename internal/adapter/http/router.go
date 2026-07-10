package httpadapter

import (
	"simple_api/internal/adapter/http/handler"
	"simple_api/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router, uc usecase.BookUseCase) {
	h := handler.NewBookHandler(uc)
	books := r.Group("/book")
	books.Get("/:id", h.GetById)
	books.Post("", h.Create)
	books.Put("/:id", h.Update)
	books.Delete("/:id", h.Delete)
}
