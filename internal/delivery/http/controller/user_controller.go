package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/wahyunurdian26/cst_app_new/internal/helper"
	"github.com/wahyunurdian26/cst_app_new/internal/model"
	"github.com/wahyunurdian26/cst_app_new/internal/usecase"
)

type UserController struct {
	UserService usecase.UserService
}

func NewUserController(userService usecase.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	request := new(model.UserCreateRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		})
	}

	response, err := c.UserService.Create(ctx.Context(), request)
	if err != nil {
		return helper.ErrorResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (h *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.UserService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Code:   200,
		Status: "OK",
		Data:   users,
	})
}

func (h *UserController) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid user ID",
		})
	}

	user, err := h.UserService.GetById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(model.Response{
			Code:   404,
			Status: "NOT FOUND",
			Data:   "User not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Code:   200,
		Status: "OK",
		Data:   user,
	})
}

func (u *UserController) Update(ctx *fiber.Ctx) error {

	var request model.UserUpdateRequest
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.Response{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid user ID",
		})

	}
	request.Id = uint(id)
	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	response, err := u.UserService.Update(ctx.Context(), &request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *UserController) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Invalid user ID",
		})
	}

	if err := h.UserService.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Code:   200,
		Status: "OK",
		Data:   "User deleted successfully",
	})
}
