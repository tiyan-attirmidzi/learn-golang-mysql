package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"api.pustaka/book"
	"api.pustaka/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	fmt.Println(dsn)
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
	v1.PATCH("/books/:id", bookHandler.Update)
	v1.DELETE("/books/:id", bookHandler.Destroy)

	router.Run()

}
