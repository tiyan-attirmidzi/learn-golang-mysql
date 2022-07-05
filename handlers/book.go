package handlers

import (
	"fmt"
	"net/http"

	"api.pustaka/book"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// with Query Params
func ExampleGetBookWithQueryParams(ctx *gin.Context) {

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

	var productInput book.ProductInput

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
