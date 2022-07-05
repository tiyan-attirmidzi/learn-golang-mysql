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

	// bookRequest := book.BookRequest{
	// 	Title: "Golang Fundamental",
	// 	Price: "180000",
	// 	// Rating:      4,
	// 	// Description: "Golang is Fast Language",
	// }

	// // bookRepository.Store(book)
	// bookService.Store(bookRequest)

	// Create

	// book := book.Book{}
	// book.Title = "JavaScript Fundamental"
	// book.Price = 140000
	// book.Rating = 5
	// book.Description = "JavaScript is Beautiful Language"

	// err = db.Create(&book).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error Creating Book Record")
	// 	fmt.Println("==========================")
	// }

	// Read

	// var book book.Book
	// var books []book.Book

	// // err = db.Debug().First(&book).Error
	// // err = db.Find(&books).Error
	// err = db.Where("rating = ?", 5).Find(&books).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error Reading Book Record")
	// 	fmt.Println("==========================")
	// }

	// // fmt.Println("Title: ", book.Title)
	// fmt.Println("Book Object: ", books)

	// Update

	// var book book.Book

	// err = db.Where("id = ?", 1).First(&book).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error Updating Book Record")
	// 	fmt.Println("==========================")
	// }

	// book.Title = "Golang nih BROOWW"

	// err = db.Save(&book).Error

	// if err != nil {
	// 	fmt.Println("Error Updating Book Record")
	// }

	// Delete

	// var book book.Book

	// err = db.Where("id = ?", 1).First(&book).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("Error Deleting Book Record")
	// 	fmt.Println("==========================")
	// }

	// book.Title = "Golang nih BROOWW"

	// err = db.Delete(&book).Error

	// if err != nil {
	// 	fmt.Println("Error Deleting Book Record")
	// }

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
	v1.POST("/books", handlers.ExamplePostBook)
	// v1.POST("/tests", handlers.Store)

	router.Run(":8000")

}
