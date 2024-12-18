package variant

import (
	"database/sql"
	"fmt"
	"practice6/model"

	"github.com/google/uuid"
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
	stmt, err := store.db.Prepare("INSERT INTO variants(id, productID, color, vsize, vprice, stock_price) VALUES (?,?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var resultVariants []model.Variant
	for _, variant := range variants {
		var addedVar model.Variant
		_, resErr := stmt.Exec(variant.ID, variant.ProductID, variant.Color, variant.Size, variant.Price, variant.Stock)
		if resErr != nil {
			fmt.Println(resErr)
			return nil, resErr
		}
		addedVar = model.Variant{ID: variant.ID, ProductID: variant.ProductID, Color: variant.Color, Size: variant.Size, Price: variant.Price, Stock: variant.Stock}
		resultVariants = append(resultVariants, addedVar)
	}
	return resultVariants, nil
}

func (store Store) Delete(variantIDs uuid.UUID, productID uuid.UUID) error {

	return nil
}
