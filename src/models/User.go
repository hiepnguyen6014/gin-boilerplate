package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `pg:"email" json:"email" gorm:"unique"`
	Name           string `pg:"name" json:"name"`
	Picture        string `pg:"picture" json:"picture"`
	Locale         string `pg:"locale" json:"locale"`
	Verified_email bool   `pg:"verified_email" json:"verified_email"`
	Id_token       string `pg:"id_token" json:"id_token"`
}
