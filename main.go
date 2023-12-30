package main

import (
	"os"
	"strings"

	"gin-boilerplate/src/modules/auth"
	"gin-boilerplate/src/modules/products"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	var (
		PORT            string   = os.Getenv("PORT")
		TRUSTED_PROXIES []string = strings.Split(os.Getenv("TRUSTED_PROXIES"), ",")
	)

	router := gin.Default()

	defer router.Run(":" + PORT)

	router.SetTrustedProxies(TRUSTED_PROXIES)

	auth.AuthRoute(router)
	products.ProductsRoute(router)

}
