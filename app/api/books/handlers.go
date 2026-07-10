package books

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"simple_api/app/dto"
	"simple_api/app/service"
)

type Handler struct {
	svc service.Book
}

func NewHandler(db *gorm.DB) *Handler {
	svc := service.NewBook(db)
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) GetById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	res, err := h.svc.GetBookById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "book not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "internal server error"})
	}
	response := Response{
		Id:     res.Id,
		Title:  res.Title,
		Author: res.Author,
		Rating: res.Rating,
	}
	return c.JSON(response)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	var body dto.BookUpdateRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	body.Id = id
	res, err := h.svc.UpdateBook(body)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "book not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "iternal server error"})
	}
	response := Response{
		Id:     res.Id,
		Title:  res.Title,
		Author: res.Author,
		Rating: res.Rating,
	}
	return c.JSON(response)

}

func (h *Handler) Create(c *fiber.Ctx) error {
	var body dto.BookCreateRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	if _, err := h.svc.CreateBook(body); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "iternal server error"})
	}
	return c.Status(201).JSON(fiber.Map{"message": "created"})
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.svc.DeleteBook(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "book not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "iternal server error"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "deleted"})
}
