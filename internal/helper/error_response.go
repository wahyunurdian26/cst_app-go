package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wahyunurdian26/cst_app_new/internal/model"
)

func ErrorResponse(ctx *fiber.Ctx, statusCode int, message string) error {
	return ctx.Status(statusCode).JSON(model.ErrorResponse{
		Code:    statusCode,
		Message: message,
	})
}
