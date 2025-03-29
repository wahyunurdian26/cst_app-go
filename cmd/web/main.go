package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wahyunurdian26/cst_app_new/db/migrations"
	"github.com/wahyunurdian26/cst_app_new/internal/config"
	"github.com/wahyunurdian26/cst_app_new/internal/delivery/http/controller"
	"github.com/wahyunurdian26/cst_app_new/internal/delivery/http/route"
	"github.com/wahyunurdian26/cst_app_new/internal/repository"
	"github.com/wahyunurdian26/cst_app_new/internal/usecase"
)

func main() {
	// Inisialisasi Fiber
	app := fiber.New()

	// Load konfigurasi
	viperConfig := config.NewViper()
	db := config.NewDatabase(viperConfig)

	// Migrasi database
	migrations.MigrateDB(db)

	// Inisialisasi Repository, Usecase, dan Controller
	userRepo := repository.NewUserRepository(db)
	userService := usecase.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Setup Routing
	routeConfig := route.RouteConfig{
		App:            app,
		UserController: userController,
	}
	routeConfig.Setup() // Memanggil fungsi untuk menambahkan routing ke Fiber

	// Jalankan server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
