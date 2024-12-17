package product

import (
	"fmt"
	"practice6/model"

	"github.com/google/uuid"
)

type Service struct {
	vsvc VariantService
	str  Store
}

func New(vsvc VariantService, str Store) Service {
	return Service{
		vsvc: vsvc,
		str:  str,
	}
}

func (service Service) Create(product *model.Product) (*model.Product, error) {
	if product.Name == "" || len(product.Name) == 0 {
		return nil, fmt.Errorf("name cannot be empty")
	}

	product.ID = uuid.New()

	product, err := service.str.Create(product)
	if err != nil {
		return nil, err
	}

	variants, err := service.vsvc.Create(product.Variants, product.ID)
	if err != nil {
		return nil, err
	}

	product.Variants = variants
	return product, nil
}
