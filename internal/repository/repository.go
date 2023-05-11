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

func (repo *Repository) AddItem(name, description string, price float64, userId uint) models.Item {
	item := models.Item{
		Name:        name,
		Price:       price,
		Description: description,
		Rating:      0,
	}
	repo.db.Create(&item)
	ui := models.UserItem{
		ItemID: item.Id,
		UserID: userId,
	}
	repo.db.Create(&ui)
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
	// check if item exist
	item := repo.GetItem(itemId)
	if item.Id == 0 {
		return models.Item{}, errors.New("item not found")
	}

	// check if user exist
	_, err := repo.GetUserByID(userId)
	if err != nil {
		return models.Item{}, err
	}

	// check if user rated item
	var r models.Rating
	repo.db.Where("user_id = ? AND item_id = ?", userId, itemId).First(&r)
	if r.ID != 0 {
		return item, errors.New("user already rated this item")
	}

	r = models.Rating{ // create rating model
		UserID: userId,
		ItemID: itemId,
		Rating: rating,
		Item:   item,
	}
	if err = repo.db.Create(&r).Error; err != nil { // if user not rated yet
		return models.Item{}, err
	}

	var ratings []models.Rating

	repo.db.Where("item_id", itemId).Find(&ratings)

	sum := int64(0)
	for _, v := range ratings {
		sum += v.Rating
	}

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

func (repo *Repository) CommentItem(userID, itemID uint, text string) ([]models.Comment, error) {
	item := repo.GetItem(itemID)
	if item.Id == 0 {
		return nil, errors.New("item not found")
	}
	_, err := repo.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	comment := models.Comment{
		Comment:   text,
		CreatedAt: time.Now(),
		UserID:    userID,
		Item:      item,
		ItemID:    itemID,
	}
	if err = repo.db.Create(&comment).Error; err != nil {
		return nil, err
	}
	// return all comments of this item
	var itemComments []models.Comment
	repo.db.Where("item_id = ?", itemID).Find(&itemComments)

	return itemComments, nil
}

func (repo *Repository) AddOrder(userID uint, itemIDs []uint) (any, error) {
	var items []models.Item
	repo.db.Where("id IN ?", itemIDs).Find(&items)

	_, err := repo.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	order := models.Order{
		UserID: userID,
		Items:  []models.OrderItem{},
		Total:  0,
		Status: true,
	}
	repo.db.Create(&order)

	for _, item := range items {
		orderItem := models.OrderItem{
			OrderID: order.ID,
			ItemID:  item.Id,
			Item:    item,
		}
		order.Total += item.Price
		repo.db.Create(&orderItem)
	}
	repo.db.Save(&order)
	var orderItems []models.OrderItem
	repo.db.Preload("Item").Find(&orderItems, "order_id = ?", order.ID)
	return orderItems, nil
}

func (repo *Repository) PurchaseItem(orderId uint) (interface{}, error) {
	var order models.Order
	repo.db.First(&order, orderId)
	if order.ID == 0 {
		return nil, nil
	}
	if order.Status == false {
		return nil, errors.New("order is false")
	}
	var bank models.Bank
	repo.db.First(&bank, order.UserID)
	if bank.UserID == 0 {
		return nil, errors.New("do not have bank account")
	}

	if order.Total > bank.Money {
		return nil, errors.New("not enough money")
	}

	pay := models.Payment{
		OrderID: orderId,
		Amount:  order.Total,
		Order:   order,
	}
	order.Status = false
	repo.db.Save(&order)
	if err := repo.db.Create(&pay).Error; err != nil {
		return nil, err
	}

	bank.Money -= order.Total
	repo.db.Save(&bank)
	return pay, nil
}

func (repo *Repository) UserBank(userId uint, money float64) bool {
	var user models.User
	if result := repo.db.First(&user, userId); result.Error != nil {
		// does not exist
		return false
	}
	bank := models.Bank{
		UserID: userId,
		Money:  money,
	}

	repo.db.Create(&bank)

	return true
}
