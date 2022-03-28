package cli

import (
	"enigmacamp.com/goacc/delivery/util"
	"enigmacamp.com/goacc/usecase"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
)

func Login(useCase usecase.AuthenticationUseCase) {
	var userName string
	var userPassword string
	fmt.Print(util.UserNameInput)
	fmt.Scanln(&userName)
	fmt.Print(util.PasswordInput)
	password, err := terminal.ReadPassword(0)
	if err == nil {
		userPassword = string(password)
	}
	if err := useCase.Login(userName, userPassword); err != nil {
		ExitApp()
	} else {
		fmt.Println("")
		MainMenu()
	}
}
