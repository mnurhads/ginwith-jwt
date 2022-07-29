package database

import (
	"ginwith-jwt/models"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

// koneksi database
func Connect(connectionString string) () {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot coonect DB")
	}

	log.Println("Connected to local Database... Very Good!")
}

// auto migrate
func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Succesfully.")
}