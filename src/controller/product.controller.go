package controller

import (
	"mime/multipart"
	"mono-golang/src/service"
	"mono-golang/src/types"
	"mono-golang/src/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddProductRequestBody struct {
	Name     string                `json:"name" form:"name"`
	Price    int                   `json:"price" form:"price"`
	Category int                   `json:"category" form:"category"`
	Quantity int                   `json:"quantity" form:"quantity"`
	Picture  *multipart.FileHeader `json:"picture" form:"picture"`
}

func GetAllProducts(ctx *gin.Context) {

	products := service.GetAllProducts()

	ctx.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func CreateProduct(ctx *gin.Context) {

	body := AddProductRequestBody{}

	if err := ctx.Bind(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pictureName, err := util.SaveUploadFile(body.Picture)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newProduct := types.AddProductBody{
		Name:     body.Name,
		Price:    body.Price,
		Quantity: body.Quantity,
		Category: body.Category,
		Picture:  pictureName,
	}

	httpCode, obj := service.CreateProduct(newProduct)

	ctx.JSON(httpCode, obj)

}
