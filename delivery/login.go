package delivery

import (
	"enigmacamp.com/goacc/delivery/util"
	"enigmacamp.com/goacc/model"
	"enigmacamp.com/goacc/repository"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
)

func Login(repo repository.UserCredentialRepo) {
	var userName string
	var userPassword string
	fmt.Print(util.UserNameInput)
	fmt.Scanln(&userName)
	fmt.Print(util.PasswordInput)
	password, err := terminal.ReadPassword(0)
	if err == nil {
		userPassword = string(password)
	}
	if repo.GetByUserNameAndPassword(model.NewUserCredential(userName, userPassword)) {
		fmt.Println("")
		MainMenu()
	} else {
		ExitApp()
	}
}
