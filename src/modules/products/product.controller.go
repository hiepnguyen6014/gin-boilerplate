package products

import "github.com/gin-gonic/gin"

func ProductController(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "ProductController",
	})
}

/*
type AddProductRequestBody struct {
	Name     string                `json:"name" form:"name"`
	Price    int                   `json:"price" form:"price"`
	Category int                   `json:"category" form:"category"`
	Quantity int                   `json:"quantity" form:"quantity"`
	Picture  *multipart.FileHeader `json:"picture" form:"picture"`
}

type UpdateProductRequestBody struct {
	Id       int    `json:"id" form:"id,default=0"`
	Name     string `json:"name" form:"name"`
	Price    int    `json:"price" form:"price,default=-1"`
	Category int    `json:"category" form:"category,default=-1"`
	Quantity int    `json:"quantity" form:"quantity,default=-1"`
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

func UpdateProduct(ctx *gin.Context) {

	body := UpdateProductRequestBody{}

	if err := ctx.Bind(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if body.Id == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "missing product id",
		})
		return
	}

	updateProduct := types.UpdateProductBody{
		Id:       body.Id,
		Name:     body.Name,
		Price:    body.Price,
		Quantity: body.Quantity,
		Category: body.Category,
	}

	httpCode, obj := service.UpdateProduct(updateProduct)

	ctx.JSON(httpCode, obj)
} */
