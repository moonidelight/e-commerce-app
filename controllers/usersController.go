package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/internal/repository"
	"project/internal/usecase"
	"project/tokens"
)

var (
	repo = repository.New()
	uc   = usecase.New(repo)
)

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			UserName string `json:"username" binding:"required"`
			Email    string `json:"email address" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		// Get the email/password off req body
		if c.Bind(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Failed to read body",
			})
			return
		}
		err := uc.SignUp(body.UserName, body.Password, body.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, "Successfully signed up")

	}
}

func LogIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			Email    string
			Password string
		}
		if c.Bind(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Failed to read body",
			})
			return
		}
		token, err := uc.Login(body.Email, body.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

func CurrentUser(c *gin.Context) {
	userID, err := tokens.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := repo.GetUserByID(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
