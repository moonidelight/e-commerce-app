package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/controllers"
	"project/internal/usecase"
	"project/middleware"
)

type Routes struct {
	uc *usecase.UseCase
}

func UserRoutes(route *gin.Engine) {
	route.POST("/signup", controllers.SignUp())
	route.POST("/login", controllers.LogIn())
	route.POST("/user/items", controllers.AddItem())
	route.GET("/user/items", controllers.GetItem())
	route.PUT("/user/items", controllers.RateItem())
	route.GET("/user/search", controllers.SearchItem())
	route.GET("/user/filtered_items", controllers.FilterItem())

	protected := route.Group("/api")
	protected.Use(middleware.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	route.POST("/logout", func(c *gin.Context) {
		c.JSON(http.StatusContinue, gin.H{})
	})
	route.PUT("/user/item/comment", controllers.CommentItem())
	route.POST("user/order", controllers.PurchaseItem())
}
