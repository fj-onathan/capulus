package database

import (
	"capulus/config/settings"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var Db *gorm.DB

// Setup initializes the database instance
func Setup() (*gorm.DB, error) {
	var err error
	Db, err = gorm.Open(setting.DatabaseSetting.Type, setting.DatabaseSetting.Local)

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
		return &gorm.DB{}, nil
	}

	Db.SingularTable(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	return Db, nil
}
