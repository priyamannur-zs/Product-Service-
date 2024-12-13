package model

import "github.com/google/uuid"

type Variant struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"productID"`
	Color     string    `json:"color"`
	Size      string    `json:"size"`
	Price     float32   `json:"price"`
	Stock     float64   `json:"stock"`
}
