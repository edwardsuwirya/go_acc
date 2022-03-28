package cli

import (
	"enigmacamp.com/goacc/delivery/util"
	"enigmacamp.com/goacc/usecase"
	"fmt"
)

func ListProductForm(useCase usecase.SearchProductUseCase) {
	util.CreateHeader(util.ListProductFormHeader)
	util.CreateHeaderTable()
	for idx, product := range useCase.Search("") {
		fmt.Printf(util.ProductListTableFormat, idx+1, product.ProductCode, product.ProductName)
	}
	ExitToMain()
}
