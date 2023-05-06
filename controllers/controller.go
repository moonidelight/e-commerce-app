package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		var body struct {
			rating int64
		}
		if c.BindJSON(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Failed to read body",
			})
			return
		}
		itemIdQuery := c.Query("id")
		itemID, err := strconv.Atoi(itemIdQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Failed to convert to ing",
			})
			return
		}
		//userID := c.Query("user_id")
		c.JSON(http.StatusOK, uc.RateItem(uint(itemID), body.rating))

	}
}
func SearchItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		item, err := uc.SearchItem(name)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err,
			})
			return
		}
		c.JSON(http.StatusFound, item)
	}
}
func FilterItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			price  float64
			rating int64
		}
		if c.Bind(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Failed to read body",
			})
			return
		}
		c.JSON(http.StatusOK, uc.FilterItemByPriceAndRating(body.price, body.rating))
	}
}
