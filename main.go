package main

import (
	"enigmacamp.com/goacc/delivery"
	"enigmacamp.com/goacc/manager"
	"fmt"
)

func main() {
	repoManager := manager.NewRepoManager()
	useCaseManager := manager.NewUseCaseManager(repoManager)

	delivery.Login(useCaseManager.UserLoginUseCase())

	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			delivery.NewProductForm(useCaseManager.ProductRegistrationUseCase())
			break
		case "2":
			delivery.ListProductForm(useCaseManager.SearchProductUseCase())
			break
		case "3":
			delivery.SearchProductForm(useCaseManager.SearchProductUseCase())
			break
		case "4":
			delivery.ExitApp()
		}
	}
}
