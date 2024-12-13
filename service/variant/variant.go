package variant

import (
	"fmt"
	"practice6/model"
)

type Service struct {
	store Store
}

func New(store Store) Service {
	return Service{
		store: store,
	}
}

func (s Service) Create(variants []model.Variant) ([]model.Variant, error) {
	for _, variant := range variants {
		if variant.Price < 0 {
			return nil, fmt.Errorf("Price cannot be negative")
		}
	}
	resVariants, err := s.store.Create(variants)
	if err != nil {
		return nil, fmt.Errorf("Some SQL Connection Error")
	}

	return resVariants, nil
}
