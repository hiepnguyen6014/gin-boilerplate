package routes

import (
	"mono-golang/src/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.GET("/google", controllers.GoogleCallback)
		auth.GET("/google/login", controllers.RedirectLoginPage)

		auth.GET("/facebook", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "facebook login",
			})
		})
		auth.GET("/facebook/login", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "facebook",
			})
		})
	}

}
