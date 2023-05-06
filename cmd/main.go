package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"project/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8181"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	//router.Use(middleware.Authentication())

	log.Fatal(router.Run(":" + port))
}
