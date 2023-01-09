package config

import (
	"fmt"
	"log"
	"mono-golang/src/constant"
	"mono-golang/src/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var CONNECTION = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
	constant.DB_HOST, constant.DB_PORT, constant.DB_USERNAME, constant.DB_PASSWORD, constant.DB_NAME)

var db *gorm.DB

func Connect() *gorm.DB {

	db, err := gorm.Open(postgres.Open(CONNECTION), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to the Database")
		panic(err)
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Product{})

	return db
}

func DB() *gorm.DB {

	log.Print(db)

	if db != nil {
		return db
	}

	return Connect()
}
