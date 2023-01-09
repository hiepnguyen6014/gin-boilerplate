package util

import (
	"log"
	"mono-golang/src/constant"
	"mono-golang/src/types"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateAccessToken(user types.User) (access_token string) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   user.Email,
		"name":    user.Name,
		"picture": user.Picture,
		"role":    user.Role,
	})

	access_token, err := token.SignedString([]byte(constant.JWT_SECRET))

	if err != nil {
		log.Print(err)
		panic(err)
	}

	return access_token
}
