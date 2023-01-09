package controller

import (
	"mono-golang/src/service"
	"mono-golang/src/types"
	"mono-golang/src/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RedirectGoogleLoginPage(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, service.RedirectGoogleLoginPage())
}

func GoogleCallback(ctx *gin.Context) {
	queries := ctx.Request.URL.Query()

	httpCode, obj := service.GoogleCallback(queries)

	googleResponse, ok := obj.(types.GoogleResponse)

	if ok {
		access_token := service.SaveUserProfile(googleResponse)

		ctx.SetCookie("access_token", access_token, 3600*24*365, "", "", false, true)
		ctx.Redirect(http.StatusMovedPermanently, "http://localhost:8080")
	}

	ctx.JSON(httpCode, obj)

}

func Authentication(ctx *gin.Context) {
	access_token, err := ctx.Cookie("access_token")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	user, _ := util.VerifyToken(access_token)

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": access_token,
		"user":         user,
	})
}
