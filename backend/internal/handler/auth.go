package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wnmay/songhub/backend/internal/entities"
	"github.com/wnmay/songhub/backend/internal/usecase"
)

type AuthHandler struct {
	authUseCase usecase.AuthUsecase
}

func NewAuthHandler (useCase usecase.AuthUsecase) *AuthHandler{
	return &AuthHandler{authUseCase: useCase}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.authUseCase.Register(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusCreated)
}