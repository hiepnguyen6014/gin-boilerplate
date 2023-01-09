package model

import "gorm.io/gorm"

type Product struct {
	Name     string `pg:"name" json:"name"`
	Price    int    `pg:"price" json:"price"`
	Category int    `pg:"category" json:"category"`
	Picture  string `pg:"picture" json:"picture"`
	Quantity int    `pg:"quantity" json:"quantity" gorm:"default:1"`
	gorm.Model
}
