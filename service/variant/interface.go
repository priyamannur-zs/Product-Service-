package variant

import "practice6/model"

type Store interface {
	Create([]model.Variant) ([]model.Variant, error)
}
