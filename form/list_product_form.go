package form

import (
	"enigmacamp.com/goacc/repository"
	"fmt"
)

func ListProductForm() {
	var mainMenuConfirmation string
	headerList()
	headerTable()
	for idx, product := range repository.GetAll() {
		fmt.Printf("%-3d%-20s%-20s\n", idx+1, product["productCode"], product["productName"])
	}
	fmt.Printf("Main Menu (y/n) ?")
	fmt.Scanln(&mainMenuConfirmation)
	if mainMenuConfirmation == "y" {
		MainMenu()
	}
}

func headerList() {
	fmt.Println("Halaman Daftar Product")
	fmt.Println("**********************")
}

func headerTable() {
	fmt.Printf("%3s%-20s%-20s\n", "No.", "Kode Produk", "Nama Produk")
}
