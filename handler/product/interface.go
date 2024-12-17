package product

import "practice6/model"

type ProSvcInf interface {
	Create(*model.Product) (*model.Product, error)
}
