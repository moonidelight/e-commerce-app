package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			name        string
			description string
			price       float64
		}
		if c.Bind(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Failed to read body",
			})
			return
		}
		items := uc.AddItem(body.name, body.description, body.price)
		c.JSON(http.StatusCreated, items)
	}
}
func GetItem() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
func RateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		//h := c.Request.Header["Authorization"]
		//var body struct {
		//	rating int64
		//}
		//itemID := c.Query("id")
		//userID := c.Query("user_id")

	}
}
func SearchItem() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
func FilterItem() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
