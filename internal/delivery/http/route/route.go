package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wahyunurdian26/cst_app_new/internal/delivery/http/controller"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *controller.UserController
}

func (c *RouteConfig) Setup() {
	c.UserRoutes()
}

func (c *RouteConfig) UserRoutes() {
	// Endpoint CRUD User
	c.App.Post("/api/users", c.UserController.CreateUser)
	c.App.Get("/api/users", c.UserController.GetAllUsers)
	c.App.Get("/api/users/:id", c.UserController.GetUserByID)
	c.App.Patch("/api/users/:id", c.UserController.Update)
	c.App.Delete("/api/users/:id", c.UserController.DeleteUser)
}
