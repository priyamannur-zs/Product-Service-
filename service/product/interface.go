package product

import "practice6/model"

type VariantService interface {
	Create([]model.Variant) ([]model.Variant, error)
}
