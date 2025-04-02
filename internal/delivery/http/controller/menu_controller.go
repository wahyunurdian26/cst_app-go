package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/wahyunurdian26/cst_app_new/internal/entity"
	"gorm.io/gorm"
)

type MenuController struct {
	DB *gorm.DB
}

func NewMenuController(db *gorm.DB) *MenuController {
	return &MenuController{DB: db}
}

// GetMenus menampilkan menu berdasarkan role dan user_email
func (m *MenuController) GetMenus(c *fiber.Ctx) error {
	// Ambil token JWT dari header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	tokenString := authHeader[len("Bearer "):] // Ambil token tanpa "Bearer "
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("XYYZ"), nil // Secret key JWT
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid or expired token"})
	}

	// Ambil email dan role dari token
	userEmail, _ := claims["email"].(string)
	userRole, _ := claims["role"].(string)

	var menus []entity.Menu

	// Ambil menu berdasarkan role
	if err := m.DB.Table("menus").
		Select("menus.*").
		Joins("JOIN role_menus rm ON rm.menu_id = menus.id").
		Where("rm.role_id = ?", userRole).
		Find(&menus).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error fetching menus by role"})
	}

	// Ambil menu berdasarkan user_email
	var userMenus []entity.Menu
	if err := m.DB.Table("menus").
		Select("menus.*").
		Joins("JOIN user_menus um ON um.menu_id = menus.id").
		Where("um.user_email = ?", userEmail).
		Find(&userMenus).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error fetching menus by user"})
	}

	// Gabungkan menu dari role_menus dan user_menus
	menus = append(menus, userMenus...)

	return c.JSON(fiber.Map{"menus": menus})
}
