package model

import "github.com/google/uuid"

type Product struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Descrption string    `json:"description"`
	Variants   []Variant `json:"variants"`
}
