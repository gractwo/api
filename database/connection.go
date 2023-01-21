package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func CreateConnection() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(
		&Quote{},
		&QuoteLine{},
		&Image{},
		&Admincard{},
		&Splash{},
		&User{},
		&Badge{},
		&GivenBadges{},
	)
}
