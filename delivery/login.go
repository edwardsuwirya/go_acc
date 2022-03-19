package delivery

import (
	"enigmacamp.com/goacc/delivery/util"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
)

func Login() {
	var userName string
	var userPassword string
	fmt.Print(util.UserNameInput)
	fmt.Scanln(&userName)
	fmt.Print(util.PasswordInput)
	password, err := terminal.ReadPassword(0)
	if err == nil {
		userPassword = string(password)
	}
	if userName == "admin" && userPassword == "123" {
		fmt.Println("")
		MainMenu()
	} else {
		ExitApp()
	}
}
