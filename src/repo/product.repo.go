package repo

import (
	"mono-golang/src/model"
	"mono-golang/src/types"
	"net/http"
)

func GetAllProducts() (products []model.Product) {

	db := db()

	db.Find(&products)

	return products
}

func CreateProduct(product types.AddProductBody) (int, interface{}) {

	newProduct := model.Product{
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
		Category: product.Category,
		Picture:  product.Picture,
	}

	db := db()

	if result := db.Create(&newProduct); result.Error != nil {
		return http.StatusBadRequest, model.Product{}
	}

	return http.StatusOK, newProduct
}
