package main

import (
	"enigmacamp.com/goacc/delivery"
	"enigmacamp.com/goacc/repository"
	"fmt"
)

func main() {
	delivery.Login()
	repo := repository.NewProductRepo()
	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			delivery.NewProductForm(repo)
			break
		case "2":
			delivery.ListProductForm(repo)
			break
		case "3":
			delivery.SearchProductForm(repo)
			break
		case "4":
			delivery.ExitApp()
		}
	}
}
