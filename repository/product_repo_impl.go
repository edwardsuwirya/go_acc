package repository

import "enigmacamp.com/goacc/model"

type productRepoImpl struct {
	productDb *[]model.Product
}

func (p *productRepoImpl) Insert(newProduct model.Product) {
	*p.productDb = append(*p.productDb, newProduct)
}

func (p *productRepoImpl) GetByProductCode(productCode string) model.Product {
	var selectedProduct model.Product
	for _, product := range *p.productDb {
		pc := product.GetProductCode()
		if pc == productCode {
			selectedProduct = product
			break
		}
	}
	return selectedProduct
}

func (p *productRepoImpl) GetAll() []model.Product {
	return *p.productDb
}

func NewProductRepo(productDb *[]model.Product) ProductRepo {
	productRepo := productRepoImpl{
		productDb: productDb,
	}
	return &productRepo
}
