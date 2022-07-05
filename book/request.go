package book

import (
	"encoding/json"
)

type BookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Rating      json.Number `json:"rating" binding:"required,number"`
	// SubTitle string      `json:"sub_title" binding:"required"` // directive
}
