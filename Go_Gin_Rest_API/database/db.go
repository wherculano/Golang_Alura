package database

import (
	"Alura/Go_Gin_Rest_API/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	strConn := "host=localhost user=root dbname=root password=root"
	DB, err := gorm.Open(postgres.Open(strConn))
	if err != nil {
		log.Panic("Connection failed")
	}
	DB.AutoMigrate(&models.Student{})
}
