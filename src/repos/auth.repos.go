package repos

import (
	"mono-golang/src/configs"
	"mono-golang/src/models"

	"gorm.io/gorm"
)

func db() *gorm.DB {
	return configs.DB()
}

func CreateUser(user models.User) (id int) {
	db := db()

	var temp models.User

	if err := db.Where("email = ?", user.Email).First(&temp).Error; err == nil {
		return int(temp.ID)
	}

	db.Create(&user)

	id = int(user.ID)

	return id
}
