package utils

func GenerateAccessToken() (access_token string) {

	/* token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
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

	return access_token */
	return "access_token"
}
