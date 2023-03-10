package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"user-rest-api/models"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	connectionString := "host=localhost user=root password=root dbname=root port=5432"
	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("Erro de conexão com o banco")
	}
	DB.AutoMigrate(&models.User{})
}
