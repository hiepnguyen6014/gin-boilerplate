package repo

import (
	"errors"
	"mono-golang/src/config"
	"mono-golang/src/model"
	"mono-golang/src/types"
	"mono-golang/src/util"

	"gorm.io/gorm"
)

func db() *gorm.DB {
	return config.DB()
}

func CreateUser(user model.User) (access_token string) {
	db := db()

	token_user := types.User{
		Name:    user.Name,
		Email:   user.Email,
		Picture: user.Picture,
		Role:    user.Role,
	}

	access_token = util.GenerateAccessToken(token_user)

	if err := db.Where("email = ?", user.Email).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		user.Access_token = access_token
		db.Create(&user)
	} else {
		user.Access_token = access_token
		db.Save(&user)
	}

	return access_token
}
