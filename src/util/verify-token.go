package util

import (
	"fmt"
	"mono-golang/src/constant"
	"mono-golang/src/types"

	"github.com/golang-jwt/jwt/v4"
)

func VerifyToken(tokenString string) (user types.User, err error) {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(constant.JWT_SECRET), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.Email = claims["email"].(string)
		user.Role = int(claims["role"].(float64))
		user.Picture = claims["picture"].(string)
		user.Name = claims["name"].(string)
	} else {
		return types.User{}, fmt.Errorf("token not valid")
	}

	return user, err
}
