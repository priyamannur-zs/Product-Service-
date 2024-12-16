package variant

import (
	"database/sql"
	"fmt"
	"practice6/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := uuid.New()
	productID := uuid.New()

	mockStore := NewMockStore(ctrl)

	Service := New(mockStore)

	testCases := []struct {
		name        string
		reqVariants []model.Variant
		pp          uuid.UUID
		expVariants []model.Variant
		modFunc     func([]model.Variant)
		wantErr     error
	}{
		{
			name:        "Valid",
			reqVariants: []model.Variant{{ID: id, ProductID: uuid.Nil, Color: "Black", Size: "8", Price: 80, Stock: 80}},
			pp:          productID,
			expVariants: []model.Variant{{ID: id, ProductID: productID, Color: "Black", Size: "8", Price: 80, Stock: 80}},
			modFunc: func(variants []model.Variant) {
				mockStore.EXPECT().Create(variants).Return(variants, nil)
			},
			wantErr: nil,
		},
		{
			name:        "SQL throws Error",
			reqVariants: []model.Variant{{ID: id, ProductID: uuid.Nil, Color: "color", Size: "9", Price: 80, Stock: 70}},
			pp:          productID,
			expVariants: nil,
			modFunc: func(variants []model.Variant) {
				mockStore.EXPECT().Create(variants).Return(nil, sql.ErrConnDone)
			},
			wantErr: fmt.Errorf("Some SQL Connection Error"),
		},
		{
			name:        "Price Validation",
			reqVariants: []model.Variant{{ID: id, ProductID: uuid.Nil, Color: "B", Size: "8", Price: -1, Stock: 8}},
			pp:          productID,
			expVariants: nil,
			modFunc: func(variants []model.Variant) {

			},
			wantErr: fmt.Errorf("Price cannot be negative"),
		},
	}

	for _, test := range testCases {
		test.modFunc(test.reqVariants)
		result, err := Service.Create(test.reqVariants, test.pp)

		for i := range result {
			result[i].ProductID = test.pp
			result[i].ID = id

		}

		assert.Equal(t, test.expVariants, result)
		assert.Equal(t, err, test.wantErr)
	}
}
