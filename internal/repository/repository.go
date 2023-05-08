package repository

import (
	"errors"
	"fmt"
	"log"
	"project/models"
	"strconv"
	"time"
)

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
		repo.db.Where("price >= ? AND (rating >= ? AND rating <= ?)", minPrice, minRating, maxRating).Order("price, rating asc").Find(&items)
	} else {
		repo.db.Where("(price >= ? AND price <= ?) AND (rating >= ? AND rating <= ?)", minPrice, maxPrice, minRating, maxRating).Order("price, rating asc").Find(&items)
	}
	return items
}

func (repo *Repository) RateItem(itemId, userId uint, rating int64) (models.Item, error) {
	item := repo.GetItem(itemId)

	// check if user exist
	_, err := repo.GetUserByID(userId)
	if err != nil {
		return models.Item{}, err
	}

	rate := models.Rating{
		UserID: userId,
		ItemID: itemId,
		Rating: rating,
	}
	if err = repo.db.Create(&rate).Error; err != nil {
		return models.Item{}, err
	}

	var ratings []models.Rating
	var sum int64
	sum = 0
	if err = repo.db.Where("item_id = ?", itemId).Find(&ratings).Error; err != nil {
		return models.Item{}, err
	}
	fmt.Println(len(ratings))
	for r := range ratings {
		sum += ratings[r].Rating
	}
	fmt.Println(sum)
	l := int64(len(ratings))
	if l == 0 {
		item.Rating = float64(rating)
	} else {
		r := float64(sum) / float64(l)
		item.Rating, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", r), 64)
	}
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

func (repo *Repository) CommentItem(userID, itemID uint, text string) (interface{}, error) {
	item := repo.GetItem(itemID)
	if item.Id == 0 {
		return nil, errors.New("item not found")
	}
	_, err := repo.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	comment := models.Comment{
		UserID:    userID,
		ItemID:    itemID,
		Comment:   text,
		CreatedAt: time.Now(),
	}
	if err = repo.db.Create(&comment).Error; err != nil {
		return nil, err
	}
	var comments []models.Comment
	repo.db.Where("item_id", itemID).Find(&comments)
	return comments, nil
}

func (repo *Repository) AddOrder(userID uint, itemIDs []uint) (any, error) {
	var items []models.Item
	repo.db.Where("id IN ?", itemIDs).Find(&items)
	fmt.Println(items)
	_, err := repo.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	order := models.Order{
		UserID:    userID,
		CreatedAt: time.Now(),
		Status:    true,
	}
	repo.db.Create(&order)

	for _, item := range items {
		orderItem := models.OrderItem{
			OrderID: order.ID,
			ItemID:  item.Id,
		}
		repo.db.Create(&orderItem)

	}
	var orders models.Order
	if err = repo.db.Preload("Items").First(&orders, order.ID).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (repo *Repository) PurchaseItem(userId, orderId uint) (interface{}, error) {
	var order models.Order
	repo.db.First(&order, orderId)
	if order.ID == 0 {
		return nil, nil
	}
	var total float64
	total = 0
	var orders models.Order
	if err := repo.db.Preload("Items").First(&orders, order.ID).Error; err != nil {
		return nil, err
	}
	itemIDs := make([]uint, len(orders.Items))
	for i, oi := range orders.Items {
		itemIDs[i] = oi.ItemID
	}
	var items []models.Item
	repo.db.Where("id in ?", itemIDs).Find(&items)
	for _, item := range items {
		total += item.Price
	}
	pay := models.Payment{
		OrderID:     orderId,
		Amount:      total,
		PaymentDate: time.Now(),
	}
	order.Status = false
	repo.db.Save(&order)
	if err := repo.db.Create(&pay).Error; err != nil {
		return nil, err
	}
	return pay, nil
}
