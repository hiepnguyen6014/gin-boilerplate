package main

import (
	"mono-golang/src/configs"
	"mono-golang/src/routes"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	// configs.SetUp()
	configs.Connect()

	routes.InitRoutes()
}
