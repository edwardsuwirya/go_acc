package delivery

import (
	"enigmacamp.com/goacc/delivery/util"
	"enigmacamp.com/goacc/repository"
	"fmt"
)

func SearchProductForm(repo repository.ProductRepo) {

	var productCode string
	util.CreateHeader(util.SearchProductFormHeader)
	util.CreateHeaderTable()
	fmt.Print(util.ProductCodeInput)
	fmt.Scanln(&productCode)
	product := repo.GetByProductCode(productCode)
	fmt.Printf(util.ProductListTableFormat, 1, product.GetProductCode(), product.GetProductName())
	ExitToMain()
}
