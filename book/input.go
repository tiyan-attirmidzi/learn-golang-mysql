package book

import "encoding/json"

type ProductInput struct {
	Name     string      `json:"name" binding:"required"`
	Price    json.Number `json:"price" binding:"required,number"`
	SubTitle string      `json:"sub_title" binding:"required"` // directive
}
