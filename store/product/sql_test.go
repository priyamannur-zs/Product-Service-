package product

import (
	"database/sql"
	"fmt"
	"practice6/model"
	"regexp"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/data-dog/go-sqlmock.v1"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	store := New(db)

	prodID := uuid.New()

	testCases := []struct {
		reqProduct *model.Product
		expProduct *model.Product
		mockFunc   func(*model.Product)
		wantErr    error
	}{
		{
			reqProduct: &model.Product{ID: prodID, Name: "IDK", Descrption: "IDKE", Variants: []model.Variant{}},
			expProduct: &model.Product{ID: prodID, Name: "IDK", Descrption: "IDKE", Variants: []model.Variant{}},
			mockFunc: func(prod *model.Product) {
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO products (id,name,description ) VALUES ( ?, ?, ?)")).
					WithArgs(prod.ID, prod.Name, prod.Descrption).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: nil,
		},

		{
			reqProduct: &model.Product{ID: prodID, Name: "IDK", Descrption: "IDKE", Variants: []model.Variant{}},
			expProduct: nil,
			mockFunc: func(prod *model.Product) {
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO products (id,name,description ) VALUES ( ?, ?, ?)")).
					WithArgs(prod.ID, prod.Name, prod.Descrption).WillReturnError(sql.ErrConnDone)

			},
			wantErr: fmt.Errorf("Some SQL error"),
		},
		{
			reqProduct: &model.Product{ID: prodID, Name: "IDK", Descrption: "IDKE", Variants: []model.Variant{}},
			expProduct: nil,
			mockFunc: func(prod *model.Product) {
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO products (id,name,description ) VALUES ( ?, ?, ?)")).
					WithArgs(prod.ID, prod.Name, prod.Descrption).WillReturnResult(sqlmock.NewResult(0, 0))
			},
			wantErr: fmt.Errorf("No rows were affected"),
		},
		{
			reqProduct: &model.Product{ID: prodID, Name: "IDK", Descrption: "IDKE", Variants: []model.Variant{}},
			expProduct: nil,
			mockFunc: func(prod *model.Product) {
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO products (id,name,description ) VALUES ( ?, ?, ?)")).
					WithArgs(prod.ID, prod.Name, prod.Descrption).WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("result Error")))
			},
			wantErr: fmt.Errorf("result Error"),
		},
	}

	for _, test := range testCases {
		test.mockFunc(test.reqProduct)
		resProduct, resErr := store.Create(test.reqProduct)

		assert.Equal(t, test.expProduct, resProduct)
		assert.Equal(t, resErr, test.wantErr)
	}
}
