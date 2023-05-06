package usecase

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"project/internal/repository"
	"project/models"
	"project/tokens"
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

func (uc *UseCase) AddItem(name, description string, price float64) models.Item {
	return uc.repo.AddItem(name, description, price)
}

func (uc *UseCase) SearchItem(name string) (models.Item, error) {
	item := uc.repo.SearchItem(name)
	if item.Id == 0 {
		return item, errors.New("item with such name doesn't exists")
	}
	return item, nil
}

func (uc *UseCase) RateItem(id uint, rate int64) models.Item {
	return uc.repo.RateItem(id, rate)
}

func (uc *UseCase) FilterItemByPriceAndRating(price float64, rating int64) []models.Item {
	return uc.repo.FilterItemByRatingAndPrice(rating, price)
}

func HashPassword(password string) (string, error) {
	const costFactor = 12

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), costFactor)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(userPassword, givenPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(givenPassword))
	if err != nil {
		return fmt.Errorf("invalid password %s", err)
	}
	return nil
}
