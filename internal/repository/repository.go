package repository

import (
	"errors"
	"fmt"
	"log"
	"project/models"
)

//var (
//	ErrInvalid = errors.New("error")
//)

func (repo *Repository) CreateUser(username, password, email string) bool {
	var user models.User
	//status := false
	if result := repo.db.Where("email = ?", email).First(&user); result.Error == nil {
		fmt.Println("------------user exists-------------")
		return false
	}

	newUser := models.User{
		UserName: username,
		Password: password,
		Email:    email,
	}
	if result := repo.db.Create(&newUser); result.Error != nil {
		fmt.Println("----------not created-----------", result.Error)
		return false
	}
	fmt.Println("----------created-------------")
	return true

}
func (repo *Repository) GetUserByID(id uint) (models.User, error) {
	var user models.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return user, errors.New("User not found!")
	}
	return user, nil
}
func (repo *Repository) Login(email string) models.User {
	var user models.User
	repo.db.First(&user, "email = ?", email)
	return user
}

func (repo *Repository) AddItem(name, description string, price float64) models.Item {
	item := models.Item{
		Name:        name,
		Price:       price,
		Description: description,
		Rating:      0,
		IsActive:    true,
	}
	repo.db.Create(&item)
	return item
}

func (repo *Repository) SearchItem(name string) models.Item {
	var item models.Item
	repo.db.Where("name = ?", name).First(&item)
	return item
}

func (repo *Repository) FilterItemByRatingAndPrice(rating int64, price float64) []models.Item {
	var items []models.Item
	repo.db.Where("price >= ? AND rating >= ?", price, rating).Order("price, rating asc").Find(&items)
	return items
}

func (repo *Repository) RateItem(id uint, rating int64) models.Item {
	item := repo.GetItem(id)
	item.Rating = (item.Rating + rating) / 2
	repo.db.Save(&item)
	return item
}

func (repo *Repository) GetItem(id uint) models.Item {
	var item models.Item
	err := repo.db.First(&item, "id = ?", id)
	if err != nil {
		log.Panic(err)
	}
	return item
}
