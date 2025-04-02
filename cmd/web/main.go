package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wahyunurdian26/cst_app_new/db/migrations"
	"github.com/wahyunurdian26/cst_app_new/internal/config"
	"github.com/wahyunurdian26/cst_app_new/internal/delivery/http/controller"
	"github.com/wahyunurdian26/cst_app_new/internal/delivery/http/route"
	"github.com/wahyunurdian26/cst_app_new/internal/repository"
	"github.com/wahyunurdian26/cst_app_new/internal/service"
)

func main() {
	// Inisialisasi Fiber
	app := fiber.New()

	// Load konfigurasi
	viperConfig := config.NewViper()
	db := config.NewDatabase(viperConfig)
	log := config.NewLogger()
	validator := config.NewValidator()

	// Migrasi database
	migrations.MigrateDB(db)

	// Inisialisasi Repository, service, dan Controller
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, validator, log)
	userController := controller.NewUserController(userService)

	campaignRepo := repository.NewCampaignRepository(db)
	campaignService := service.NewCampaignService(campaignRepo, validator, log)
	campaignController := controller.NewCampaignController(campaignService)

	// Setup Routing
	routeConfig := route.RouteConfig{
		App:                app,
		UserController:     userController,
		CampaignController: campaignController,
	}
	routeConfig.Setup() // Memanggil fungsi untuk menambahkan routing ke Fiber

	// Jalankan server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
