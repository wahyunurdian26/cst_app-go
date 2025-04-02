package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wahyunurdian26/cst_app_new/internal/delivery/http/controller"
	"github.com/wahyunurdian26/cst_app_new/internal/delivery/http/middleware"
	"gorm.io/gorm"
)

type RouteConfig struct {
	DB                 *gorm.DB
	App                *fiber.App
	UserController     *controller.UserController
	CampaignController *controller.CampaignController
	AuthController     *controller.AuthController
	MenuController     *controller.MenuController
}

func (c *RouteConfig) Setup() {
	c.UserRoutes()
	c.CampaignRoutes()
	c.AuthRoutes()
	c.MenuRoutes()
}

func (c *RouteConfig) UserRoutes() {
	c.App.Post("/api/users", middleware.RoleMiddleware("ROL000"), c.UserController.CreateUser)
	c.App.Get("/api/users", middleware.RoleMiddleware("ROL000"), c.UserController.GetAllUsers)
	c.App.Get("/api/users/:id", middleware.RoleMiddleware("ROL000"), c.UserController.GetUserByID)
	c.App.Patch("/api/users/:id", middleware.RoleMiddleware("ROL000"), c.UserController.UpdateUser)
	c.App.Delete("/api/users/:id", middleware.RoleMiddleware("ROL000"), c.UserController.DeleteUser)
}

func (c *RouteConfig) CampaignRoutes() {
	c.App.Get("/api/campaign/offers", c.CampaignController.GetAllOffer)
	c.App.Get("/api/campaign/senders", c.CampaignController.GetAllSender)
	c.App.Get("/api/campaign/products", c.CampaignController.GetAllProduct)
	c.App.Get("/api/campaign/brands", c.CampaignController.GetAllBrand)

	c.App.Post("/api/campaign/create", c.CampaignController.CreateCampaign)
	c.App.Get("/api/campaign/campaigns", c.CampaignController.GetAllCampaign)
	c.App.Get("/api/campaign/:id_campaign", c.CampaignController.GetCampaignByID)

}

func (c *RouteConfig) AuthRoutes() {

	c.App.Post("/api/auth/login", c.AuthController.Login)

}

func (c *RouteConfig) MenuRoutes() {
	menuGroup := c.App.Group("/api/menu")

	// Endpoint untuk menampilkan semua menu yang dapat diakses berdasarkan role user
	menuGroup.Get("/", middleware.RoleMiddleware("ROL000", "ROL001", "ROL002"), c.MenuController.GetMenus)
}
