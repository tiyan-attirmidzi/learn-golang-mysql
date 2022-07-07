package book

import (
	"encoding/json"
)

type BookRequestStore struct {
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Rating      json.Number `json:"rating" binding:"required,number"`
	// SubTitle string      `json:"sub_title" binding:"required"` // directive
}

type BookRequestUpdate struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Price       json.Number `json:"price" binding:"number"`
	Rating      json.Number `json:"rating" binding:"number"`
}
