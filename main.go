package main

import (
	"os"

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
	db.AutoMigrate(&Quote{},
		&QuoteLine{})

	//Init api
	router := InitRouter()
	err = router.Run(":2021")

}
