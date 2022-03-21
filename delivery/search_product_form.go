package delivery

import (
	"enigmacamp.com/goacc/delivery/util"
	"enigmacamp.com/goacc/usecase"
	"fmt"
)

func SearchProductForm(useCase usecase.SearchProductUseCase) {

	var productCode string
	util.CreateHeader(util.SearchProductFormHeader)
	fmt.Print(util.ProductCodeInput)
	fmt.Scanln(&productCode)
	product := useCase.Search(productCode)
	if product != nil {
		util.CreateHeaderTable()
		p := product[0]
		fmt.Printf(util.ProductListTableFormat, 1, p.GetProductCode(), p.GetProductName())
	} else {
		fmt.Println("Produk tidak ditemukan")
	}

	ExitToMain()
}
