package repository

import (
	"enigmacamp.com/goacc/model"
	"github.com/jmoiron/sqlx"
)

type productRepoImpl struct {
	productDb *sqlx.DB
}

func (p *productRepoImpl) Insert(newProduct model.Product) {
	//*p.productDb = append(*p.productDb, newProduct)
}

func (p *productRepoImpl) GetByProductCode(productCode string) model.Product {
	var selectedProduct model.Product
	//for _, product := range *p.productDb {
	//	pc := product.GetProductCode()
	//	if pc == productCode {
	//		selectedProduct = product
	//		break
	//	}
	//}
	return selectedProduct
}

func (p *productRepoImpl) GetAll() []model.Product {
	//return *p.productDb
	return nil
}

func NewProductRepo(productDb *sqlx.DB) ProductRepo {
	productRepo := productRepoImpl{
		productDb: productDb,
	}
	return &productRepo
}
