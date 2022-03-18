package form

import (
	"enigmacamp.com/goacc/repository"
	"fmt"
)

func NewProductForm() {
	var productCode string
	var productName string
	var saveProductConfirmation string
	headerAdd()
	fmt.Print("Kode Produk : ")
	fmt.Scanln(&productCode)
	fmt.Print("Nama Produk : ")
	fmt.Scanln(&productName)
	fmt.Printf("Produk %s dengan kode %s akan disimpan (y/n) ?", productName, productCode)
	fmt.Scanln(&saveProductConfirmation)

	if saveProductConfirmation == "y" {
		repository.Insert(productCode, productName)
	}
	MainMenu()
}

func headerAdd() {
	fmt.Println("Halaman Pendaftaran Produk")
	fmt.Println("**************************")
}
