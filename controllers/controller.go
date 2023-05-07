package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			Name        string
			Description string
			Price       float64
		}
		if c.Bind(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Failed to read body",
			})
			return
		}
		items := uc.AddItem(body.Name, body.Description, body.Price)
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
			Rating int64
		}
		if c.BindJSON(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Failed to read body",
			})
			return
		}
		itemIdQuery := c.Query("item_id")
		itemID, err := strconv.Atoi(itemIdQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Failed to convert to int",
			})
			return
		}
		userIdQuery := c.Query("user_id")
		userID, err := strconv.Atoi(userIdQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Failed to convert to int",
			})
			return
		}

		rate, err := uc.RateItem(itemID, userID, body.Rating)
		c.JSON(http.StatusOK, rate)

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
			MinPrice  float64 `json:"min_price" default:"0"`
			MaxPrice  float64 `json:"max_price" default:"-1"`
			MinRating int64   `json:"min_rating" default:"0"`
			MaxRating int64   `json:"max_rating" default:"5"`
		}
		if c.Bind(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Failed to read body",
			})
			return
		}

		c.JSON(http.StatusOK, uc.FilterItemByPriceAndRating(body.MinRating, body.MaxRating, body.MinPrice, body.MaxPrice))
	}
}

func CommentItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			Comment string
		}
		if c.Bind(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Failed to read body",
			})
			return
		}
		itemID, err := strconv.Atoi(c.Query("item_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Failed to convert to int",
			})
			return
		}
		userID, err := strconv.Atoi(c.Query("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Failed to convert to int",
			})
			return
		}
		item, err := uc.CommentItem(userID, itemID, body.Comment)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err,
			})
			return
		}
		c.JSON(http.StatusCreated, item)
	}
}

func PurchaseItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		itemID, _ := strconv.Atoi(c.Query("item_id"))
		userID, _ := strconv.Atoi(c.Query("user_id"))
		orders, err := uc.PurchaseItem(userID, itemID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err,
			})
			return
		}
		c.JSON(http.StatusCreated, orders)
	}
}
