package delivery

import (
	"enigmacamp.com/goacc/delivery/util"
	"enigmacamp.com/goacc/usecase"
	"fmt"
)

func NewProductForm(usecase usecase.ProductRegistrationUseCase) {
	var productCode string
	var productName string
	var saveProductConfirmation string
	util.CreateHeader(util.NewProductFormHeader)
	fmt.Print(util.ProductCodeInput)
	fmt.Scanln(&productCode)
	fmt.Print(util.ProductNameInput)
	fmt.Scanln(&productName)
	fmt.Printf(util.SaveProductConfirmation, productName, productCode)
	fmt.Scanln(&saveProductConfirmation)

	if saveProductConfirmation == "y" {
		usecase.Register(productCode, productName)
	}
	MainMenu()
}
