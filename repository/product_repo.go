package repository

import "enigmacamp.com/goacc/model"

type ProductRepo interface {
	Insert(newProduct model.Product)
	GetByProductCode(productCode string) model.Product
	GetAll() []model.Product
}
