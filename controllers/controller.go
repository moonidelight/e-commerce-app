package controllers

import (
	"fmt"
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
func GetItemList() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
func RateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
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
			MinPrice  float64 `json:"min_price"`
			MaxPrice  float64 `json:"max_price"`
			MinRating int64   `json:"min_rating"`
			MaxRating int64   `json:"max_rating"`
		}
		if c.BindJSON(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Failed to read body",
			})
			return
		}

		if body.MaxRating == 0 {
			body.MaxRating = 5
		}
		if body.MaxPrice == 0 {
			body.MaxPrice = -1
		}

		fmt.Println(body.MinPrice, body.MaxPrice, body.MinRating, body.MaxRating)

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

func AddOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := strconv.Atoi(c.Query("user_id"))
		var body struct {
			Items []int
		}
		if c.BindJSON(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Failed to read body",
			})
			return
		}
		orders, err := uc.AddOrder(userID, body.Items)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err,
			})
			return
		}
		c.JSON(http.StatusCreated, orders)
	}
}

func PurchaseItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Query("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err,
			})
			return
		}
		orderId, err := strconv.Atoi(c.Query("order_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err,
			})
			return
		}
		info, err := uc.PurchaseItem(userId, orderId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err,
			})
			return
		}
		c.JSON(http.StatusOK, info)
	}
}
