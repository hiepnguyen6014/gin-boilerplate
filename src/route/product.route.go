package route

import (
	"mono-golang/src/controller"

	"github.com/gin-gonic/gin"
)

func ProductRoute(router *gin.Engine) {
	product := router.Group("/product")
	{
		product.POST("", controller.CreateProduct)
		product.GET("", controller.GetAllProducts)
	}
}
