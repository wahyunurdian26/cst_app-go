package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/wahyunurdian26/cst_app_new/internal/entity"
	"gorm.io/gorm"
)

// MenuAccessMiddleware memeriksa apakah pengguna memiliki akses ke menu tertentu
func MenuAccessMiddleware(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil token JWT dari header Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		tokenString := strings.Split(authHeader, " ")
		if len(tokenString) != 2 || tokenString[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token format",
			})
		}

		// Parsing token JWT
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString[1], claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("XYYZ"), nil // Gantilah ini dengan secret key yang sesuai
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid or expired token",
			})
		}

		// Ambil email pengguna dari token
		userEmail, ok := claims["email"].(string)
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Email not found in token",
			})
		}

		// Ambil menu yang diminta dari parameter URL
		menuID := c.Params("menu_id")
		if menuID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Menu ID is required",
			})
		}

		// Periksa apakah pengguna memiliki akses ke menu ini
		var userMenu entity.UserMenu
		if err := db.Where("user_email = ? AND menu_id = ?", userEmail, menuID).First(&userMenu).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"message": "Access denied to this menu",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Database error",
			})
		}

		// Jika akses diizinkan, lanjutkan ke handler berikutnya
		return c.Next()
	}
}
