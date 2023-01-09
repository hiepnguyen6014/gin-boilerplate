package route

import (
	"mono-golang/src/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.GET("/google", controller.GoogleCallback)
		auth.GET("/google/login", controller.RedirectGoogleLoginPage)
		auth.GET("", controller.Authentication)
	}
}
