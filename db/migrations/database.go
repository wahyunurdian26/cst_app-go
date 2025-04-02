package migrations

import (
	"log"

	"github.com/wahyunurdian26/cst_app_new/internal/entity"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Group{},
		&entity.SubGroup{},
		&entity.Brand{},
		&entity.Offer{},
		&entity.Sender{},
		&entity.Product{},
		&entity.Campaign{},
		&entity.Menu{},
		&entity.RoleMenu{},
		&entity.UserMenu{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migrated successfully!")
}
