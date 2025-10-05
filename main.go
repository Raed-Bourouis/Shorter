package main

import (
	"fmt"

	"github.com/Raed-Bourouis/Shorter/handler"
	"github.com/Raed-Bourouis/Shorter/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Go URL shortener"})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortURL(c)
	})
	r.GET("/:shortURL", func(c *gin.Context) {
		handler.HandleRedirect(c)
	})

	store.InitStore()

	err := r.Run(":9000")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
