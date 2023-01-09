package route

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {

	var (
		PORT            string   = os.Getenv("PORT")
		TRUSTED_PROXIES []string = strings.Split(os.Getenv("TRUSTED_PROXIES"), ",")
	)

	router := gin.Default()

	router.Static("/assets", "./uploads")

	defer router.Run(":" + PORT)

	router.SetTrustedProxies(TRUSTED_PROXIES)

	AuthRoute(router)
	ProductRoute(router)
}
