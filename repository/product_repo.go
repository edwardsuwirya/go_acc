package repository

type Product map[string]string

var productList []Product

func Insert(productCode string, productName string) {
	newProduct := make(Product)
	newProduct["productCode"] = productCode
	newProduct["productName"] = productName
	productList = append(productList, newProduct)
}

func GetByProductCode(productCode string) string {
	var selectedProductName string
	for _, product := range productList {
		p := product["productCode"]
		if p == productCode {
			selectedProductName = p
			break
		}
	}
	return selectedProductName
}

func GetAll() []Product {
	return productList
}
