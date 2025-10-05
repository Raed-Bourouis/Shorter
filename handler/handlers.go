package handler

import (
	"fmt"
	"net/http"

	"github.com/Raed-Bourouis/Shorter/shortener"
	"github.com/Raed-Bourouis/Shorter/store"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortURL(c *gin.Context) {
	var creationRequest UrlCreationRequest

	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURL := shortener.GenerateShortURL(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveURLMapping(shortURL, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:9000/"
	c.JSON(http.StatusCreated, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortURL,
	})

}

func HandleRedirect(c *gin.Context) {
	shortURL:=c.Param("shortURL")
	longURL:=store.GetLongURL(shortURL)
	fmt.Println(longURL)
	c.Redirect(http.StatusTemporaryRedirect,longURL)

}
