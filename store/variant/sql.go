package variant

import (
	"database/sql"
	"practice6/model"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return Store{
		db: db,
	}
}

func (store Store) Create(variants []model.Variant) ([]model.Variant, error) {
	stmt, err := store.db.Prepare("INSERT INTO variants(id, productID, color, size, price, stock) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return nil, err
	}
	var resultVariants []model.Variant
	for _, variant := range variants {
		var addedVar model.Variant
		_, resErr := stmt.Exec(variant.ID, variant.ProductID, variant.Color, variant.Size, variant.Price, variant.Stock)
		if resErr != nil {
			return nil, resErr
		}
		addedVar = model.Variant{ID: variant.ID, ProductID: variant.ProductID, Color: variant.Color, Size: variant.Size, Price: variant.Price, Stock: variant.Stock}
		resultVariants = append(resultVariants, addedVar)
	}
	return resultVariants, nil
}
