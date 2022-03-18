package main

import (
	"enigmacamp.com/goacc/form"
	"fmt"
	"os"
)

func main() {
	form.MainMenu()

	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			form.NewProductForm()
			break
		case "2":
			form.ListProductForm()
			break
		case "3":
			os.Exit(0)
		}
	}
}
