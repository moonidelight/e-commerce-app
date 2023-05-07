package usecase

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

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
