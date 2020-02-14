package migrations

import (
	"capulus/database/migrations/app"
	"github.com/jinzhu/gorm"
	"log"
)

// Migrate automigrates models using ORM
func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&app.Host{},
	)

	// Success
	log.Print("Auto Migration has beed processed with")
}
