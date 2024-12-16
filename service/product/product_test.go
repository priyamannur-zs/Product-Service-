package product

import (
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

	store := NewMockStore(ctrl)
	vservice := NewMockVariantService(ctrl)

	pservice := New(vservice, store)

	testCases := []struct {
		reqProduct *model.Product
		expProduct *model.Product
		mockFunc   func(*model.Product)
		wantErr    error
	}{
		{
			reqProduct: &model.Product{Name: "name", Descrption: "description"},
			expProduct: &model.Product{ID: uuid.Nil, Name: "name", Descrption: "description"},
			mockFunc: func(p *model.Product) {
				vservice.EXPECT().Create(p.Variants, gomock.Any()).Return(p.Variants, nil)
				// (func(variants []model.Variant, pid uuid.UUID) ([]model.Variant, error) {
				// 	for i := range variants {
				// 		variants[i].ID = uuid.Nil
				// 		variants[i].ProductID = uuid.Nil
				// 	}
				// 	return p.Variants, nil
				store.EXPECT().Create(p).DoAndReturn(func(product *model.Product) (*model.Product, error) {
					product.ID = uuid.Nil
					return product, nil
				})
			},
			wantErr: nil,
		},
		{
			reqProduct: &model.Product{ID: uuid.Nil, Name: "name", Descrption: "description"},
			expProduct: nil,
			mockFunc: func(p *model.Product) {
				vservice.EXPECT().Create(p.Variants, gomock.Any()).Return(p.Variants, nil)
				// DoAndReturn(func(variants []model.Variant, pid uuid.UUID) ([]model.Variant, error) {
				// 	// for i := range variants {
				// 	// 	variants[i].ID = uuid.Nil
				// 	// 	variants[i].ProductID = uuid.Nil
				// 	// }
				// 	return variants, nil
				// })
				store.EXPECT().Create(p).Return(nil, fmt.Errorf("Some SQL error"))

			},
			wantErr: fmt.Errorf("Some SQL error"),
		},
		{
			reqProduct: &model.Product{ID: uuid.Nil, Name: "name", Descrption: "description"},
			expProduct: nil,
			mockFunc: func(p *model.Product) {
				vservice.EXPECT().Create(p.Variants, gomock.Any()).Return(nil, fmt.Errorf("No rows were affected"))
				//DoAndReturn(func(variants []model.Variant, pid uuid.UUID) ([]model.Variant, error) {
				// for i := range variants {
				// 	variants[i].ID = uuid.Nil
				// 	variants[i].ProductID = uuid.Nil
				// }
				//return nil, fmt.Errorf("No rows were affected")
				//})
				// store.EXPECT().Create(p).DoAndReturn(func(product *model.Product) (*model.Product, error) {
				// 	product.ID = uuid.Nil
				// 	return product, nil
				// })
			},
			wantErr: fmt.Errorf("No rows were affected"),
		},
		{
			reqProduct: &model.Product{ID: uuid.Nil, Name: "name", Descrption: "description"},
			expProduct: nil,
			mockFunc: func(p *model.Product) {
				vservice.EXPECT().Create(p.Variants, gomock.Any()).DoAndReturn(func(variants []model.Variant, pid uuid.UUID) ([]model.Variant, error) {
					// for i := range variants {
					// 	variants[i].ID = uuid.Nil
					// 	variants[i].ProductID = uuid.Nil
					// }
					return nil, fmt.Errorf("result Error")
				})
				// store.EXPECT().Create(p).DoAndReturn(func(product *model.Product) (*model.Product, error) {
				// 	product.ID = uuid.Nil
				// 	return product, nil
				// })
			},
			wantErr: fmt.Errorf("result Error"),
		},
		{
			reqProduct: &model.Product{ID: uuid.Nil, Name: "name", Descrption: "description"},
			expProduct: nil,
			mockFunc: func(p *model.Product) {
				vservice.EXPECT().Create(p.Variants, gomock.Any()).Return(nil, fmt.Errorf("price cannot be negative"))
				//DoAndReturn(func(variants []model.Variant, pid uuid.UUID) ([]model.Variant, error) {
				// for i := range variants {
				// 	variants[i].ID = uuid.Nil
				// 	variants[i].ProductID = uuid.Nil
				// }
				//	return nil, fmt.Errorf("price cannot be negative")
				//})
				// store.EXPECT().Create(p).DoAndReturn(func(product *model.Product) (*model.Product, error) {
				// 	product.ID = uuid.Nil
				// 	return product, nil
				// })
			},
			wantErr: fmt.Errorf("price cannot be negative"),
		},
		{
			reqProduct: &model.Product{ID: uuid.Nil, Name: "", Descrption: "description"},
			expProduct: nil,
			mockFunc: func(p *model.Product) {
				// vservice.EXPECT().Create(p.Variants, gomock.Any()).DoAndReturn(func(variants []model.Variant, pid uuid.UUID) ([]model.Variant, error) {
				// 	for i := range variants {
				// 		variants[i].ID = uuid.Nil
				// 		variants[i].ProductID = uuid.Nil
				// 	}
				// 	return variants, nil
				// })
				// store.EXPECT().Create(p).DoAndReturn(func(product *model.Product) (*model.Product, error) {
				// 	product.ID = uuid.Nil
				// 	return product, nil
				// })
			},
			wantErr: fmt.Errorf("name cannot be empty"),
		},
	}

	for _, test := range testCases {
		test.mockFunc(test.reqProduct)
		resProduct, resErr := pservice.Create(test.reqProduct)

		assert.Equal(t, resProduct, test.expProduct)
		assert.Equal(t, resErr, test.wantErr)

	}
}
