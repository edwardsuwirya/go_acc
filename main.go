package main

import (
	"enigmacamp.com/goacc/delivery"
	"enigmacamp.com/goacc/manager"
	"fmt"
)

func main() {
	repoManager := manager.NewRepoManager()
	authRepo := repoManager.UserCredentialRepo()
	productRepo := repoManager.ProductRepo()

	delivery.Login(authRepo)
	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			delivery.NewProductForm(productRepo)
			break
		case "2":
			delivery.ListProductForm(productRepo)
			break
		case "3":
			delivery.SearchProductForm(productRepo)
			break
		case "4":
			delivery.ExitApp()
		}
	}
}
