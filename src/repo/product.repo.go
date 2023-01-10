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

func UpdateProduct(product types.UpdateProductBody) (int, interface{}) {
	updateProduct := model.Product{}
	db := db()

	if result := db.First(&updateProduct, product.Id); result.Error != nil {
		return http.StatusBadRequest, model.Product{}
	}

	if product.Name != "" {
		updateProduct.Name = product.Name
	}

	if product.Category >= 0 {
		updateProduct.Category = product.Category
	}

	if product.Quantity >= 0 {
		updateProduct.Quantity = product.Quantity
	}

	if product.Price >= 0 {
		updateProduct.Price = product.Price
	}

	db.Save(&updateProduct)

	return http.StatusOK, updateProduct
}
