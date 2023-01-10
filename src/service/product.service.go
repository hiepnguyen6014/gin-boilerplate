package service

import (
	"mono-golang/src/model"
	"mono-golang/src/repo"
	"mono-golang/src/types"
)

func GetAllProducts() (products []model.Product) {

	products = repo.GetAllProducts()

	return products
}

func CreateProduct(product types.AddProductBody) (httpCode int, obj interface{}) {

	httpCode, obj = repo.CreateProduct(product)

	return httpCode, obj
}

func UpdateProduct(product types.UpdateProductBody) (httpCode int, obj interface{}) {
	httpCode, obj = repo.UpdateProduct(product)

	return httpCode, obj
}
