package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wahyunurdian26/cst_app_new/internal/delivery/http/controller"
)

type RouteConfig struct {
	App                *fiber.App
	UserController     *controller.UserController
	CampaignController *controller.CampaignController
	AuthController     *controller.AuthController
}

func (c *RouteConfig) Setup() {
	c.UserRoutes()
	c.CampaignRoutes()
	c.AuthRoutes()
}

func (c *RouteConfig) UserRoutes() {
	c.App.Post("/api/users", c.UserController.CreateUser)
	c.App.Get("/api/users", c.UserController.GetAllUsers)
	c.App.Get("/api/users/:id", c.UserController.GetUserByID)
	c.App.Patch("/api/users/:id", c.UserController.UpdateUser)
	c.App.Delete("/api/users/:id", c.UserController.DeleteUser)
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
