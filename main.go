package main

import (
	"log"
	"net/http"

	"api.pustaka/book"
	"api.pustaka/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:18M@reT20@tcp(127.0.0.1:3306)/api_pustaka?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Error Connection")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handlers.NewBookHandler(bookService)

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

	v1.GET("/books", bookHandler.Index)
	v1.GET("/books/:id", bookHandler.Show)
	v1.POST("/books", bookHandler.Store)

	router.Run()

}
