package product

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"practice6/model"
)

type Handler struct {
	psvc ProSvcInf
}

func New(psvc ProSvcInf) Handler {
	return Handler{
		psvc: psvc,
	}
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {

	dataRead, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var product model.Product
	err = json.Unmarshal(dataRead, &product)
	if err != nil {
		fmt.Println(err)
		return
	}

	creProduct := &product
	resProduct, creErr := h.psvc.Create(creProduct)

	if creErr != nil {
		w.WriteHeader(400)
		_, writeErr := w.Write([]byte(creErr.Error()))
		if writeErr != nil {
			fmt.Println(writeErr)
		}
		return
	}

	mBody, marshalErr := json.Marshal(resProduct)
	if marshalErr != nil {
		fmt.Println(marshalErr)
		return
	}

	w.WriteHeader(201)

	_, writeErr := w.Write(mBody)
	if writeErr != nil {
		fmt.Println(writeErr)
	}
}
