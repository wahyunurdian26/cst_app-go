package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wahyunurdian26/cst_app_new/internal/helper"
	"github.com/wahyunurdian26/cst_app_new/internal/model"
	"github.com/wahyunurdian26/cst_app_new/internal/service"
)

type AuthController struct {
	AuthService service.AuthService
}

// Constructor untuk AuthController
func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

// Handler untuk login
func (h *AuthController) Login(c *fiber.Ctx) error {
	if len(c.Body()) == 0 {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Request body is required")
	}

	var req model.LoginUserRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Validasi login
	token, err := h.AuthService.Login(req)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
	}

	// Jika berhasil, kirimkan token dalam response
	return helper.JSONResponse(c, fiber.StatusOK, "Login successful", fiber.Map{
		"token": token,
	})
}
