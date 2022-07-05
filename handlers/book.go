package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"api.pustaka/book"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (bookHandler *bookHandler) Index(ctx *gin.Context) {

	books, err := bookHandler.bookService.FindAll()

	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Books Retrived Successfully",
		"data":    books,
	})

}

func (bookHandler *bookHandler) Show(ctx *gin.Context) {

	// TODO: Add Error Handling

	id, _ := strconv.Atoi(ctx.Param("id"))

	book, err := bookHandler.bookService.FindByID(id)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Books Retrived Successfully",
		"data":    book,
	})

}

// with Query Params
func (bookHandler *bookHandler) ExampleGetBookWithQueryParams(ctx *gin.Context) {

	name := ctx.Query("name")

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Products Retrived Successfully",
		"data": map[string]interface{}{
			"id":    4,
			"name":  name,
			"price": 20000,
		},
	})
}

// with Path Variable
func ExampleGetBookWithPathVariable(ctx *gin.Context) {

	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product Retrived Successfully",
		"data": map[string]interface{}{
			"id":    id,
			"name":  "Rokok Surya 12",
			"price": 20000,
		},
	})
}

func ExamplePostBook(ctx *gin.Context) {

	var productInput book.BookRequest

	err := ctx.ShouldBindJSON(&productInput)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errorMessages,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product Stored Successfully",
		"data":    productInput,
	})
}

// func Store(ctx *gin.Context) {

// 	var book book.Book

// 	err := ctx.ShouldBindJSON(&book)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	err = db.
// 	fmt.Println(book)

// }
