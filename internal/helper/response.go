package helper

import "github.com/gofiber/fiber/v2"

// StatusText mengembalikan teks status berdasarkan kode HTTP
func StatusText(code int) string {
	switch code {
	case fiber.StatusOK:
		return "OK"
	case fiber.StatusCreated:
		return "CREATED"
	case fiber.StatusBadRequest:
		return "BAD REQUEST"
	case fiber.StatusUnauthorized:
		return "UNAUTHORIZED"
	case fiber.StatusForbidden:
		return "FORBIDDEN"
	case fiber.StatusNotFound:
		return "NOT FOUND"
	case fiber.StatusInternalServerError:
		return "INTERNAL SERVER ERROR"
	default:
		return "UNKNOWN STATUS"
	}
}

// JSONResponse mengembalikan response standar JSON
func JSONResponse(c *fiber.Ctx, code int, message string, data interface{}) error {
	return c.Status(code).JSON(map[string]interface{}{
		"code":    code,
		"status":  StatusText(code),
		"message": message,
		"data":    data,
	})
}

// ErrorResponse mengembalikan response error standar JSON
func ErrorResponse(c *fiber.Ctx, code int, message string) error {
	return JSONResponse(c, code, message, nil)
}
