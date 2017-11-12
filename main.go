package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	// Web API definiton here
	router.GET("/quote", getQuote)
	router.POST("/quote", getQuote)

	router.GET("/video", getVideoURL)
	router.POST("/video", getVideoURL)

	router.Run(":" + port)
}
