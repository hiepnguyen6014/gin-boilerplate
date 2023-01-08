package configs

import (
	"fmt"
	"log"
	"mono-golang/src/constants"
	"mono-golang/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var CONNECTION = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
	constants.DB_HOST, constants.DB_PORT, constants.DB_USERNAME, constants.DB_PASSWORD, constants.DB_NAME)

var db *gorm.DB

func Connect() *gorm.DB {

	db, err := gorm.Open(postgres.Open(CONNECTION), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to the Database")
		panic(err)
	}

	db.AutoMigrate(&models.User{})

	log.Print("Connected DB!")

	return db
}

func DB() *gorm.DB {

	if db != nil {
		return db
	}

	return Connect()
}
