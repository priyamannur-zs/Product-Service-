package product

import (
	"database/sql"
	"fmt"
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

func (s Store) Create(product *model.Product) (*model.Product, error) {
	result, err := s.db.Exec("INSERT INTO products (id,name,description ) VALUES ( ?, ?, ?)", product.ID, product.Name, product.Descrption)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Some SQL error")
	}

	if ra, rerr := result.RowsAffected(); rerr != nil || ra <= 0 {
		if rerr != nil {
			return nil, fmt.Errorf("result Error")
		}
		return nil, fmt.Errorf("No rows were affected")
	}

	return product, nil
}
