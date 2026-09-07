package api

import (
	"fmt"
	"net/http"
	"shortlink/internal/model"
	"time"

	"github.com/gin-gonic/gin"
)

type ShortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}

func CreateShortLink(c *gin.Context) {
	var req ShortenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid link format"})
		return
	}

	shortCode, err := model.GenerateShortLink(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server internal error"})
		return
	}

	shortURL := fmt.Sprintf("http://localhost:8080/%s", shortCode)

	c.JSON(http.StatusOK, gin.H{
		"short_url": shortURL,
		"long_url":  req.URL,
	})
}

func Redirect(c *gin.Context) {
	shortCode := c.Param("short_code")

	longURL, err := model.RDB.Get(model.Ctx, "shortlink:"+shortCode).Result()

	if err == nil {
		fmt.Println("Read from redis cache")
		c.Redirect(http.StatusFound, longURL)
		return
	}

	fmt.Println("Cache didn't match, Search it in the MySQL")
	longURL, err = model.GetLongURL(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short link doesn't exist"})
		return
	}

	model.RDB.Set(model.Ctx, "shortlink:"+shortCode, longURL, 24*time.Hour)

	c.Redirect(http.StatusFound, longURL)
}
