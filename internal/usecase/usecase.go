package usecase

import (
	"errors"
	"project/internal/repository"
	"project/models"
	"project/tokens"
	"time"
)

type UseCase struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{repo: repo}
}

var (
	ErrCreateUser     = errors.New("can't create user")
	ErrHashPassword   = errors.New("can't hash password")
	ErrRecordNotFound = errors.New("such user doesn't exits")
	ErrGenerateToken  = errors.New("can't generate token")
)

func (uc *UseCase) SignUp(username, password, email string) error {
	hashedPassword, err := HashPassword(password) // HashPassword func is in the below of this code
	if err != nil {
		return ErrHashPassword
	}
	check := uc.repo.CreateUser(username, hashedPassword, email)
	if !check {
		return ErrCreateUser
	}
	return nil
}

func (uc *UseCase) Login(email, givenPassword string) (string, error) {
	user := uc.repo.Login(email)
	if user.Id == 0 {
		return "", ErrRecordNotFound
	}
	if err := VerifyPassword(user.Password, givenPassword); err != nil {
		return "", err
	}
	token, err := tokens.GenToken(user.Id)
	if err != nil {
		return "", ErrGenerateToken
	}
	return token, nil
}

func (uc *UseCase) AddItem(name, description string, price float64, userId int) models.Item {
	return uc.repo.AddItem(name, description, price, uint(userId))
}

func (uc *UseCase) SearchItem(name string) (models.Item, error) {
	item := uc.repo.SearchItem(name)
	if item.Id == 0 {
		return item, errors.New("item with such name doesn't exists")
	}
	return item, nil
}

func (uc *UseCase) RateItem(itemId, userId int, rate int64) (models.Item, error) {
	item, user := uint(itemId), uint(userId)
	return uc.repo.RateItem(item, user, rate)
}

func (uc *UseCase) FilterItemByPriceAndRating(minRating, maxRating int64, minPrice, maxPrice float64) []models.Item {
	return uc.repo.FilterItemByRatingAndPrice(minRating, maxRating, minPrice, maxPrice)
}

type responseComment struct {
	ID        uint
	Comment   string
	CreatedAt time.Time
	UserID    uint
}

func (uc *UseCase) CommentItem(userID, itemID int, text string) (interface{}, error) {
	user := uint(userID)
	item := uint(itemID)

	c, err := uc.repo.CommentItem(user, item, text)

	comments := make([]responseComment, len(c))
	for i, comment := range c {
		comments[i].ID = comment.ID
		comments[i].Comment = comment.Comment
		comments[i].CreatedAt = comment.CreatedAt
		comments[i].UserID = comment.UserID
	}
	return comments, err
}
func (uc *UseCase) AddOrder(userID int, itemIDs []int) (any, error) {
	user := uint(userID)
	var items []uint
	for _, id := range itemIDs {
		items = append(items, uint(id))
	}
	return uc.repo.AddOrder(user, items)
}
func (uc *UseCase) PurchaseItem(order int) (interface{}, error) {

	return uc.repo.PurchaseItem(uint(order))
}

func (uc *UseCase) UserBank(userId int, money float64) bool {
	return uc.repo.UserBank(uint(userId), money)
}
