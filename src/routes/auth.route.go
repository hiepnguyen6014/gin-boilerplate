package routes

import (
	"mono-golang/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	auth := router.Group("/auth")

	auth.GET("/google", func(ctx *gin.Context) {
		queries := ctx.Request.URL.Query()

		code := queries.Get("code")
		state := queries.Get("state")

		if state == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "session error",
			})
			return
		}

		data, _ := utils.GetTokenByCode(code)

		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	auth.GET("/google/login", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, utils.GetGoogleUri())
	})

	auth.GET("/facebook", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"message": "facebook login",
		})
	})

}
