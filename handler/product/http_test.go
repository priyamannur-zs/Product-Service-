package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"practice6/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocksvc := NewMockProSvcInf(ctrl)

	handler := New(mocksvc)

	testCases := []struct {
		reqProduct *model.Product
		expProduct *model.Product
		expStatus  int
		mockFunc   func(*model.Product)
		wantErr    error
	}{
		{
			reqProduct: &model.Product{ID: uuid.Nil, Name: "name", Descrption: "desc", Variants: []model.Variant{{}}},
			expProduct: &model.Product{ID: uuid.Nil, Name: "name", Descrption: "desc", Variants: []model.Variant{{}}},
			expStatus:  201,
			mockFunc: func(product *model.Product) {
				mocksvc.EXPECT().Create(product).DoAndReturn(func(product *model.Product) (*model.Product, error) {
					product.ID = uuid.Nil
					return product, nil
				})
			},
			wantErr: nil,
		},
		{
			reqProduct: &model.Product{Name: "", Descrption: "desc", Variants: []model.Variant{{}}},
			expProduct: nil,
			expStatus:  400,
			mockFunc: func(product *model.Product) {
				mocksvc.EXPECT().Create(product).Return(nil, fmt.Errorf("name cannot be empty"))
			},
			wantErr: fmt.Errorf("name cannot be empty"),
		},
	}

	for _, test := range testCases {
		test.mockFunc(test.reqProduct)

		rr := httptest.NewRecorder()

		toBytes, err := json.Marshal(test.reqProduct)
		if err != nil {
			fmt.Println(err)
			return
		}

		req, err := http.NewRequest(http.MethodPost, "/product", bytes.NewBuffer(toBytes))
		if err != nil {
			fmt.Println(err)
			return
		}

		handler.Create(rr, req)

		assert.Equal(t, test.expStatus, rr.Result().StatusCode)
	}
}
