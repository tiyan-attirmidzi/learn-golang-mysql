package main

import (
	"net/http"

	"api.pustaka/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Data Retrived Successfully",
			"data": map[string]interface{}{
				"name": "Tiyan Attirmidzi",
				"age":  22,
				"bio":  "Do It, Until U Idol Becomes U Rivals",
			},
		})
	})

	v1.GET("/books", handlers.ExampleGetBookWithQueryParams)
	v1.GET("/books/:id", handlers.ExampleGetBookWithPathVariable)
	v1.POST("/books", handlers.ExamplePostBook)

	router.Run()

}
