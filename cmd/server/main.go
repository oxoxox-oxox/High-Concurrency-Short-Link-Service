package main

import (
	"fmt"
	"shortlink/internal/api"
	"shortlink/internal/model"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Short link start running")

	dsn := "root:1@tcp(127.0.0.1:3306)/shortlink_db?charset=utf8mb4&parseTime=True&loc=Local"
	model.InitDB(dsn)

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/:shorten", api.CreateShortLink)
	}

	r.GET("/:short_code", api.Redirect)

	r.Run(":8080")
}
