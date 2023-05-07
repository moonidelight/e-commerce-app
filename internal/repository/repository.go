package repository

import (
	"errors"
	"fmt"
	"log"
	"project/models"
	"time"
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
		IsActive:    true,
		Comments:    []models.Comment{},
		Rating:      0,
	}
	repo.db.Create(&item)
	return item
}

func (repo *Repository) SearchItem(name string) models.Item {
	var item models.Item
	repo.db.Where("name = ?", name).First(&item)
	return item
}

func (repo *Repository) FilterItemByRatingAndPrice(minRating, maxRating int64, minPrice, maxPrice float64) []models.Item {
	var items []models.Item
	if maxPrice == -1 {
		repo.db.Where("price >= ? AND rating >= ? AND rating <= ?", minPrice, minRating, maxRating).Order("price, rating asc").Find(&items)
	}
	repo.db.Where("price >= ? AND price <= ? AND rating >= ? AND rating <= ?", minPrice, maxPrice, minRating, maxRating).Order("price, rating asc").Find(&items)
	return items
}

func (repo *Repository) RateItem(itemId, userId uint, rating int64) (models.Item, error) {
	item := repo.GetItem(itemId)

	// check if user exits
	_, err := repo.GetUserByID(userId)
	if err != nil {
		return models.Item{}, err
	}
	rate := models.Rating{
		UserID: userId,
		ItemID: itemId,
		Rating: rating,
	}
	repo.db.Create(&rate)

	var ratings []models.Rating
	var sum int64
	sum = 0
	repo.db.Where("item_id = ?", itemId).Find(&ratings)

	for r := range ratings {
		sum += ratings[r].Rating
	}
	item.Rating = float64(sum / int64(len(ratings)))
	repo.db.Save(&item)
	return item, nil
}

func (repo *Repository) GetItem(id uint) models.Item {
	var item models.Item
	err := repo.db.First(&item, id).Error
	if err != nil {
		log.Panic(err)
	}
	return item
}

func (repo *Repository) CommentItem(userID, itemID uint, text string) (models.Item, error) {
	item := repo.GetItem(itemID)
	if item.Id == 0 {
		return models.Item{}, errors.New("item not found")
	}
	_, err := repo.GetUserByID(userID)
	if err != nil {
		return models.Item{}, errors.New("user not found")
	}
	comment := models.Comment{
		UserID:    userID,
		Comment:   text,
		CreatedAt: time.Now(),
	}
	repo.db.Create(&comment)
	item.Comments = append(item.Comments, comment)
	if err := repo.db.Save(&item).Error; err != nil {
		return models.Item{}, errors.New("can't add comment")
	}
	return item, nil
}

func (repo *Repository) PurchaseItem(userID, itemID uint) (any, error) {
	item := repo.GetItem(itemID)
	if item.Id == 0 {
		return nil, errors.New("item not found")
	}
	_, err := repo.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	orders := models.UserOrders{
		UserID: userID,
		Status: true,
	}
	orders.Orders = append(orders.Orders, item)
	if err = repo.db.Create(&orders).Error; err != nil {
		return nil, errors.New("can't purchase item")
	}
	return orders, nil
}
