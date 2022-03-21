package delivery

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
	if useCase.Login(userName, userPassword) {
		fmt.Println("")
		MainMenu()
	} else {
		ExitApp()
	}
}
