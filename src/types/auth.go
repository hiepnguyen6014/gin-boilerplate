package types

type GoogleResponse struct {
	Access_token string `json:"access_token"`
}

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Role    int    `json:"role"`
	Picture string `json:"picture"`
}