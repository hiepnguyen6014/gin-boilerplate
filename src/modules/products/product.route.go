package products

import (
	"github.com/gin-gonic/gin"
)

func ProductsRoute(router *gin.Engine) {
	product := router.Group("/product")
	{
		product.POST("", ProductController)
		/* product.GET("", controller.GetAllProducts)
		product.PUT("", controller.UpdateProduct) */
	}
}
