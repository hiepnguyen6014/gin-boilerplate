package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.GET("/google", AuthController)
		/* auth.GET("/google/login", RedirectGoogleLoginPage)
		auth.GET("", RedirectGoogleLoginPage)
		auth.GET("/logout", RedirectGoogleLoginPage) */
	}
}
