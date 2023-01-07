package main

import (
	"mono-golang/src/routes"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	routes.InitRoutes()
}
