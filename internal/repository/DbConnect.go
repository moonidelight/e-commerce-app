package repository

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	models2 "project/models"
)

type Repository struct {
	db *gorm.DB
}

func New() *Repository {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable timezone=Asia/Almaty",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	fmt.Println(os.Getenv("DB_HOST"))
	fmt.Println(os.Getenv("DB_USER"))
	fmt.Println(os.Getenv("DB_PASSWORD"))
	fmt.Println(os.Getenv("DB_NAME"))
	//dsn := "host=localhost user=postgres password=postgres database=gorm port=5432 sslmode=disable timezone=Asia/Almaty"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect")
	}
	log.Println("Connected")
	db.AutoMigrate(&models2.User{}, &models2.Item{}, &models2.Comment{}, &models2.Rating{}, &models2.Order{}, &models2.OrderItem{}, &models2.Payment{})
	return &Repository{db: db}
}
