package variant

import (
	"fmt"
	"practice6/model"

	"github.com/google/uuid"
)

type Service struct {
	store Store
}

func New(store Store) Service {
	return Service{
		store: store,
	}
}

func (s Service) Create(variants []model.Variant, pid uuid.UUID) ([]model.Variant, error) {
	for i, variant := range variants {
		if variant.Price < 0 {
			return nil, fmt.Errorf("Price cannot be negative")
		}
		variants[i].ID = uuid.New()
		variants[i].ProductID = pid
	}
	resVariants, err := s.store.Create(variants)
	if err != nil {
		return nil, fmt.Errorf("Some SQL Connection Error")
	}

	return resVariants, nil
}
