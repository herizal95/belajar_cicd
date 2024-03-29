package config

import (
	"log"
	"os"

	"github.com/herizal95/golang_gin_starter/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("error Loading .env File")

	}
	conn := os.Getenv("Database_Connection")
	DB, err = gorm.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}

	// migrate models to database postgres

	DB.AutoMigrate(&models.Book{})

	if err != nil {
		return
	}
}
