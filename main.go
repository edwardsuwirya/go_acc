package main

import (
	"enigmacamp.com/goacc/config"
	"enigmacamp.com/goacc/delivery"
	"fmt"
)

func main() {
	appConfig := config.NewConfig()

	delivery.Login(appConfig.UseCaseManager.UserLoginUseCase())

	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			delivery.NewProductForm(appConfig.UseCaseManager.ProductRegistrationUseCase())
			break
		case "2":
			delivery.ListProductForm(appConfig.UseCaseManager.SearchProductUseCase())
			break
		case "3":
			delivery.SearchProductForm(appConfig.UseCaseManager.SearchProductUseCase())
			break
		case "4":
			delivery.ExitApp()
		}
	}
}
