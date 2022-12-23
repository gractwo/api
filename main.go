package main

import (
	"os"

	"gractwo-api/database"
	"gractwo-api/router"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//line valid only on dev. ignore
	godotenv.Load("local.env")

	//database connection
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(
		&database.Quote{},
		&database.QuoteLine{},
		&database.Image{})

	//Init api
	router := router.InitRouter()
	err = router.Run(":2021")
	if err != nil {
		panic("failed to connect database")
	}
}
