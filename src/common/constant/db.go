package constant

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

var (
	DB_HOST     = os.Getenv("DB_HOST")
	DB_PORT, _  = strconv.Atoi(os.Getenv("DB_PORT"))
	DB_NAME     = os.Getenv("DB_NAME")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
)
