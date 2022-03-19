package delivery

import (
	"enigmacamp.com/goacc/delivery/util"
	"enigmacamp.com/goacc/repository"
	"fmt"
)

func ListProductForm(repo repository.ProductRepo) {
	util.CreateHeader(util.ListProductFormHeader)
	util.CreateHeaderTable()
	for idx, product := range repo.GetAll() {
		fmt.Printf(util.ProductListTableFormat, idx+1, product.GetProductCode(), product.GetProductName())
	}
	ExitToMain()
}
