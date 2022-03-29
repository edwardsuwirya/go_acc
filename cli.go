package main

import (
	"enigmacamp.com/goacc/config"
	"enigmacamp.com/goacc/delivery/cli"
	"enigmacamp.com/goacc/manager"
	"fmt"
)

type Cli struct {
	UseCaseManager manager.UseCaseManager
}

func (c *Cli) Run() {
	cli.Login(c.UseCaseManager.UserLoginUseCase())
	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			cli.NewProductForm(c.UseCaseManager.ProductRegistrationUseCase())
			break
		case "2":
			cli.ListProductForm(c.UseCaseManager.SearchProductUseCase())
			break
		case "3":
			cli.SearchProductForm(c.UseCaseManager.SearchProductUseCase())
			break
		case "4":
			cli.ExitApp()
		}
	}
}

func Console() AppServer {
	c := config.NewConfig(".", "config")
	infraManager := manager.NewInfra(c)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	return &Cli{
		UseCaseManager: useCaseManager,
	}
}
