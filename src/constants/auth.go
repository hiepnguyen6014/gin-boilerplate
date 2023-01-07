package constants

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	GOOGLE_AUTH_USER_INFO string = os.Getenv("GOOGLE_AUTH_USER_INFO")
	GOOGLE_CLIENT_SECRET  string = os.Getenv("GOOGLE_CLIENT_SECRET")
	GOOGLE_REDIRECT_URI   string = os.Getenv("GOOGLE_REDIRECT_URI")
	GOOGLE_TOKEN_URI      string = os.Getenv("GOOGLE_TOKEN_URI")
	GOOGLE_CLIENT_ID      string = os.Getenv("GOOGLE_CLIENT_ID")
	GOOGLE_AUTH_URI       string = os.Getenv("GOOGLE_AUTH_URI")
)
