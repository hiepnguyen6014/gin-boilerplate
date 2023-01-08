package models

import "gorm.io/gorm"

type User struct {
	Username string `pg:"username" json:"username" gorm:"unique;default:null"`
	Password string `pg:"password" json:"password" gorm:"default:null"`
	Role     int    `pg:"role" json:"role" gorm:"default:0"`
	Email    string `pg:"email" json:"email" gorm:"unique"`
	Name     string `pg:"name" json:"name"`
	Picture  string `pg:"picture" json:"picture"`
	gorm.Model
	Id_token string `pg:"id_token" json:"id_token"`
}
