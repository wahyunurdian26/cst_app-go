package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wahyunurdian26/cst_app_new/internal/delivery/http/controller"
)

type RouteConfig struct {
	App                 *fiber.App
	UserController      *controller.UserController
	CamapaignController *controller.CampaignController
}

func (c *RouteConfig) Setup() {
	c.UserRoutes()
	c.CampaignRoutes()
}

func (c *RouteConfig) UserRoutes() {
	c.App.Post("/api/users", c.UserController.CreateUser)
	c.App.Get("/api/users", c.UserController.GetAllUsers)
	c.App.Get("/api/users/:id", c.UserController.GetUserByID)
	c.App.Patch("/api/users/:id", c.UserController.UpdateUser)
	c.App.Delete("/api/users/:id", c.UserController.DeleteUser)
}

func (c *RouteConfig) CampaignRoutes() {
	c.App.Get("/api/campaign/offers", c.CamapaignController.GetAllOffer)
	c.App.Get("/api/campaign/senders", c.CamapaignController.GetAllSender)
	c.App.Get("/api/campaign/products", c.CamapaignController.GetAllProduct)
	c.App.Get("/api/campaign/brands", c.CamapaignController.GetAllBrand)

	c.App.Post("/api/campaign/create", c.CamapaignController.CreateCampaign)
	c.App.Get("/api/campaign/campaigns", c.CamapaignController.GetAllCampaign)

}
