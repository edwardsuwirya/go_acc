package model

import "fmt"

type Product struct {
	productCode string
	productName string
}

func (p *Product) ToString() string {
	return fmt.Sprintf("Product Info : %s, %s", p.productCode, p.productName)
}

func (p *Product) SetProductCode(code string) {
	p.productCode = code
}
func (p *Product) SetProductName(name string) {
	p.productName = name
}
func (p *Product) GetProductCode() string {
	return p.productCode
}
func (p *Product) GetProductName() string {
	return p.productName
}
func NewProduct(productCode, productName string) Product {
	return Product{
		productCode: productCode,
		productName: productName,
	}
}
