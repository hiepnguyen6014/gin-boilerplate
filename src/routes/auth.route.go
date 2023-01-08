package routes

import (
	"mono-golang/src/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.GET("/google", controllers.GoogleCallback)
		auth.GET("/google/login", controllers.RedirectGoogleLoginPage)

	}

}
