package variant

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
		return
	}

	defer db.Close()

	id := uuid.New()
	productID := uuid.New()

	query := regexp.QuoteMeta("INSERT INTO variants(id, productID, color, size, price, stock) VALUES (?,?,?,?,?,?)")

	testCases := []struct {
		name        string
		reqVariants []model.Variant
		expVariants []model.Variant
		mockFunc    func(reqVariants []model.Variant)
		wantErr     error
	}{
		{
			name:        "Valid Testcase",
			reqVariants: []model.Variant{{ID: id, ProductID: productID, Color: "Black", Size: "3", Price: 80.0, Stock: 60}},
			expVariants: []model.Variant{{ID: id, ProductID: productID, Color: "Black", Size: "3", Price: 80.0, Stock: 60}},
			mockFunc: func(reqVariants []model.Variant) {
				mock.ExpectPrepare(query)
				for _, ereqVariant := range reqVariants {
					mock.ExpectExec(query).WithArgs(id, productID, ereqVariant.Color, ereqVariant.Size, ereqVariant.Price, ereqVariant.Stock).WillReturnResult(sqlmock.NewResult(1, 1))
				}

			},
			wantErr: nil,
		},
		{
			name:        "Invalid Testcase",
			reqVariants: []model.Variant{{ID: id, ProductID: productID, Color: "Black", Size: "3", Price: 80.0, Stock: 60}},
			expVariants: nil,
			mockFunc: func(reqVariants []model.Variant) {
				mock.ExpectPrepare(query)
				for _, ereqVariant := range reqVariants {
					mock.ExpectExec(query).WithArgs(id, productID, ereqVariant.Color, ereqVariant.Size, ereqVariant.Price, ereqVariant.Stock).WillReturnError(sql.ErrConnDone)
				}
			},
			wantErr: sql.ErrConnDone,
		},
	}

	for _, test := range testCases {
		fmt.Println(test.name)
		variantStore := New(db)

		test.mockFunc(test.reqVariants)

		result, err := variantStore.Create(test.reqVariants)

		assert.Equal(t, test.expVariants, result)
		assert.Equal(t, test.wantErr, err)
	}
}
