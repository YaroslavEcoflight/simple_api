package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"simple_api/internal/adapter/http/dto"
	"simple_api/internal/domain"
	"simple_api/internal/usecase"
)

type BookHandler struct {
	uc usecase.BookUseCase
}

func NewBookHandler(uc usecase.BookUseCase) *BookHandler {
	return &BookHandler{uc: uc}
}

func (h *BookHandler) GetById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	book, err := h.uc.GetById(id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "book not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "internal server error"})
	}
	return c.JSON(dto.BookResponse{Id: book.Id, Title: book.Title, Author: book.Author, Rating: book.Rating})
}

func (h *BookHandler) Create(c *fiber.Ctx) error {
	var body dto.BookCreateRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	if _, err := h.uc.Create(body.Title, body.Author, body.Rating); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "internal server error"})
	}
	return c.Status(201).JSON(fiber.Map{"message": "created"})
}

func (h *BookHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	var body dto.BookUpdateRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	book, err := h.uc.Update(id, body.Title, body.Author, body.Rating)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "book not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "internal server error"})
	}
	return c.JSON(dto.BookResponse{Id: book.Id, Title: book.Title, Author: book.Author, Rating: book.Rating})
}

func (h *BookHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.uc.Delete(id); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return c.Status(404).JSON(fiber.Map{"error": "book not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "internal server error"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "deleted"})
}
