package product

import (
	"practice6/model"

	"github.com/google/uuid"
)

type VariantService interface {
	Create([]model.Variant, uuid.UUID) ([]model.Variant, error)
}

type Store interface {
	Create(*model.Product) (*model.Product, error)
}
