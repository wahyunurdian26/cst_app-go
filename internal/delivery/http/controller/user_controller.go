package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/wahyunurdian26/cst_app_new/internal/entity"
	"github.com/wahyunurdian26/cst_app_new/internal/usecase"
	"github.com/wahyunurdian26/cst_app_new/internal/web"
)

type UserController struct {
	userService usecase.UserService
}

func NewUserController(userService usecase.UserService) *UserController {
	return &UserController{userService: userService}
}

func (h *UserController) CreateUser(c *fiber.Ctx) error {
	var user entity.User

	// body := c.Body()
	// println("Request Body:", string(body))

	// Parse JSON request body ke struct User
	if err := c.BodyParser(&user); err != nil {
		response := web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid request format",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Simpan user menggunakan service
	if err := h.userService.Create(&user); err != nil {
		response := web.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	// Berhasil
	response := web.WebResponse{
		Code:   201,
		Status: "CREATED",
		Data: fiber.Map{
			"id":        user.ID,
			"email":     user.Email,
			"username":  user.Username,
			"createdAt": user.CreatedAt,
		},
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAll()
	if err != nil {
		response := web.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   users,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *UserController) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response := web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid user ID",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	user, err := h.userService.GetById(uint(id))
	if err != nil {
		response := web.WebResponse{
			Code:   404,
			Status: "NOT FOUND",
			Data:   "User not found",
		}
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *UserController) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response := web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid user ID",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Ambil data JSON sebagai map
	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		response := web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid request format",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Panggil service untuk update
	if err := h.userService.Update(uint(id), updates); err != nil {
		response := web.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "User updated successfully",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *UserController) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response := web.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid user ID",
		}
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := h.userService.Delete(uint(id)); err != nil {
		response := web.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "User deleted successfully",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
