package model

import "fmt"

type Product struct {
	Id          string
	ProductCode string
	ProductName string
	CategoryId  string
}

func (p *Product) ToString() string {
	return fmt.Sprintf("Product Info : %s, %s", p.ProductCode, p.ProductName)
}

func NewProduct(productCode, productName, categoryId string) Product {
	return Product{
		ProductCode: productCode,
		ProductName: productName,
		CategoryId:  categoryId,
	}
}
