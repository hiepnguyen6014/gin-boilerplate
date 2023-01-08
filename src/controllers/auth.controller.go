package controllers

import (
	"mono-golang/src/services"
	"mono-golang/src/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RedirectGoogleLoginPage(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, services.RedirectGoogleLoginPage())
}

func GoogleCallback(ctx *gin.Context) {
	queries := ctx.Request.URL.Query()

	httpCode, obj := services.GoogleCallback(queries)

	googleResponse, ok := obj.(types.GoogleResponse)

	if ok {
		defer services.SaveUserProfile(googleResponse)

		ctx.SetCookie("access_token", googleResponse.Access_token, 3600, "", "", false, true)
		ctx.SetCookie("id_token", googleResponse.Id_token, 3600*24*30, "", "", false, true)
		ctx.Redirect(http.StatusMovedPermanently, "http://localhost:8080")
	}

	ctx.JSON(httpCode, obj)

}
