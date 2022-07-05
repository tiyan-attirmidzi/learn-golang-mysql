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

func (h *bookHandler) Index(ctx *gin.Context) {

	books, err := h.bookService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := convertToBookResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Books Retrived Successfully",
		"data":    booksResponse,
	})

}

func (h *bookHandler) Show(ctx *gin.Context) {

	// TODO: Create Error Handling (if id dont exist on db, accepted and data value is null)

	id, _ := strconv.Atoi(ctx.Param("id"))

	b, err := h.bookService.FindByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}

	bookResponse := convertToBookResponse(b)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book Retrived Successfully",
		"data":    bookResponse,
	})

}

func (h *bookHandler) Store(ctx *gin.Context) {

	var bookInput book.BookRequest

	err := ctx.ShouldBindJSON(&bookInput)

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

	book, err := h.bookService.Store(bookInput)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product Stored Successfully",
		"data":    book,
	})

}

// with Query Params
func (h *bookHandler) ExampleGetBookWithQueryParams(ctx *gin.Context) {

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

func convertToBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
	}
}
