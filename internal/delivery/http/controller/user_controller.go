package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/wahyunurdian26/cst_app_new/internal/helper"
	"github.com/wahyunurdian26/cst_app_new/internal/model"
	"github.com/wahyunurdian26/cst_app_new/internal/service"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (h *UserController) CreateUser(c *fiber.Ctx) error {
	if len(c.Body()) == 0 {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Request body is required")
	}

	request := new(model.UserCreateRequest)
	if err := c.BodyParser(request); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	response, err := h.UserService.Create(c.Context(), request)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helper.JSONResponse(c, fiber.StatusCreated, "User successfully created", response)
}

func (h *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.UserService.GetAll()
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve users")
	}

	return helper.JSONResponse(c, fiber.StatusOK, "Users retrieved successfully", users)
}

func (h *UserController) GetUserByID(c *fiber.Ctx) error {
	idStr := c.Params("id")

	// Parse string ke UUID
	id, err := uuid.Parse(idStr)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID format")
	}

	user, err := h.UserService.GetById(uuid.UUID(id))
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return helper.ErrorResponse(c, fiberErr.Code, fiberErr.Message)
		}
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Something went wrong")
	}

	if user == nil {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "User not found")
	}

	return helper.JSONResponse(c, fiber.StatusOK, "User retrieved successfully", user)
}

func (h *UserController) UpdateUser(c *fiber.Ctx) error {
	if len(c.Body()) == 0 {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Request body is required")
	}

	var request model.UserUpdateRequest
	idStr := c.Params("id")

	// Parse string ke UUID
	id, err := uuid.Parse(idStr)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID format")
	}

	if err := c.BodyParser(&request); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	request.Id = uuid.UUID(id)
	response, err := h.UserService.Update(c.Context(), &request)
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return helper.ErrorResponse(c, fiberErr.Code, fiberErr.Message)
		}
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helper.JSONResponse(c, fiber.StatusOK, "User updated successfully", response)

}

func (h *UserController) DeleteUser(c *fiber.Ctx) error {
	idStr := c.Params("id")

	// Parse string ke UUID
	id, err := uuid.Parse(idStr)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID format")
	}

	err = h.UserService.Delete(uuid.UUID(id))
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return helper.ErrorResponse(c, fiberErr.Code, fiberErr.Message)
		}
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helper.JSONResponse(c, fiber.StatusOK, "User deleted successfully", nil)
}
