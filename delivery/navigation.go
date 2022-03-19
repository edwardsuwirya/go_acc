package delivery

import (
	"enigmacamp.com/goacc/delivery/util"
	"fmt"
	"os"
)

func ExitToMain() {
	var mainMenuConfirmation string
	fmt.Printf(util.ExitToMainConfirmation)
	fmt.Scanln(&mainMenuConfirmation)
	if mainMenuConfirmation == "y" {
		MainMenu()
	}
}
func ExitApp() {
	fmt.Println("Byee...")
	os.Exit(0)
}
