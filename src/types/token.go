package types

type GoogleResponse struct {
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
	Scope        string `json:"scope"`
	Token_type   string `json:"token_type"`
	Id_token     string `json:"id_token"`
}

type GoogleUserProfile struct {
	Id             string `json:"id"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Picture        string `json:"picture"`
	Locale         string `json:"locale"`
	Verified_email bool   `json:"verified_email"`
}
