package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	models2 "project/models"
)

type Repository struct {
	db *gorm.DB
}

func New() *Repository {
	//dsn := fmt.Sprintf("host=%s user=%s password=%s db_name=%s port=5432 sslmode=disable timezone=Asia/Almaty",
	//	os.Getenv("DB_HOST"),
	//	os.Getenv("DB_USER"),
	//	os.Getenv("DB_PASSWORD"),
	//	os.Getenv("DB_NAME"),
	//)
	dsn := "host=localhost user=postgres password=postgres database=gorm port=5432 sslmode=disable timezone=Asia/Almaty"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect")
	}
	log.Println("Connected")
	db.AutoMigrate(&models2.User{}, &models2.Item{}, &models2.Comment{}, &models2.UserOrders{}, &models2.Rating{})
	return &Repository{db: db}
}
