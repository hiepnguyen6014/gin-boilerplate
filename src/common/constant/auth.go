package constant

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	GOOGLE_AUTH_USER_INFO = string(os.Getenv("GOOGLE_AUTH_USER_INFO"))
	GOOGLE_CLIENT_SECRET  = os.Getenv("GOOGLE_CLIENT_SECRET")
	GOOGLE_REDIRECT_URI   = os.Getenv("GOOGLE_REDIRECT_URI")
	GOOGLE_TOKEN_URI      = os.Getenv("GOOGLE_TOKEN_URI")
	GOOGLE_CLIENT_ID      = os.Getenv("GOOGLE_CLIENT_ID")
	GOOGLE_AUTH_URI       = os.Getenv("GOOGLE_AUTH_URI")
)

var (
	FACEBOOK_AUTH_USER_INFO = string(os.Getenv("FACEBOOK_AUTH_USER_INFO"))
	FACEBOOK_CLIENT_SECRET  = os.Getenv("FACEBOOK_CLIENT_SECRET")
	FACEBOOK_REDIRECT_URI   = os.Getenv("FACEBOOK_REDIRECT_URI")
	FACEBOOK_TOKEN_URI      = os.Getenv("FACEBOOK_TOKEN_URI")
	FACEBOOK_CLIENT_ID      = os.Getenv("FACEBOOK_CLIENT_ID")
	FACEBOOK_AUTH_URI       = os.Getenv("FACEBOOK_AUTH_URI")
)

var (
	JWT_SECRET = string(os.Getenv("JWT_SECRET"))
)