package repos

import (
	"log"
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

	if result := db.Where("email = ?", user.Email).First(&temp); result.Error != nil {
		log.Fatal(result.Error)
	}

	if temp != (models.User{}) {
		return int(temp.ID)
	}

	db.Create(&user)

	id = int(user.ID)

	return id
}
