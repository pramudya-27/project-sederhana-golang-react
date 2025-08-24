package main

import (
	"daimyo/backend-api/config"
	"daimyo/backend-api/database"
	"daimyo/backend-api/routes"
)

func main() {

	config.LoadEnv()
	database.InitDB()
	r := routes.SetupRouter()

	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
