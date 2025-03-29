package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wahyunurdian26/cst_app_new/db/migrations"
	"github.com/wahyunurdian26/cst_app_new/internal/config"
)

func main() {
	app := fiber.New()

	viperConfig := config.NewViper()
	db := config.NewDatabase(viperConfig)

	migrations.MigrateDB(db)

	app.Listen(":8080")
}
