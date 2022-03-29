package repository

import (
	"enigmacamp.com/goacc/logger"
	"enigmacamp.com/goacc/model"
	"enigmacamp.com/goacc/utils"
	"github.com/jmoiron/sqlx"
	"time"
)

type productRepoImpl struct {
	productDb *sqlx.DB
}

func (p *productRepoImpl) Insert(newProduct model.Product) (model.Product, error) {
	id := utils.GetUuid()
	newProduct.Id = id
	timestamp := time.Now()
	_, err := p.productDb.Exec("insert into m_product(id,product_code,product_name,category_id,created_at,updated_at) "+
		"values ($1,$2,$3,$4,$5,$6)",
		id, newProduct.ProductCode, newProduct.ProductName, newProduct.CategoryId, timestamp, timestamp)
	if err != nil {
		logger.Log.Error().Err(err).Str("service", "db-product-insert").Msg("Insert product failed")
		return model.Product{}, err
	}
	return newProduct, nil
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
