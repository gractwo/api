package main

import (
	"gractwo-api/database"
	"gractwo-api/router"

	"github.com/joho/godotenv"
)

func main() {
	//line valid only on dev. ignore
	godotenv.Load("local.env")

	//init database
	database.CreateConnection()

	//Init api
	router := router.InitRouter()
	err := router.Run(":2021")
	if err != nil {
		panic("failed to connect database")
	}
}
