package model

type User struct {
	Username     string `pg:"username" json:"username" gorm:"unique;default:null"`
	Password     string `pg:"password" json:"password" gorm:"default:null"`
	Role         int    `pg:"role" json:"role" gorm:"default:0"`
	Email        string `pg:"email" json:"email" gorm:"unique"`
	Name         string `pg:"name" json:"name"`
	Picture      string `pg:"picture" json:"picture"`
	Access_token string `pg:"access_token" json:"access_token"`
	// gorm.Model
}
