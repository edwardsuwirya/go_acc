package repository

import (
	"enigmacamp.com/goacc/model"
	"enigmacamp.com/goacc/utils"
)

type userCredentialRepoImpl struct {
	userCredDb *[]model.UserCredential
}

func (u *userCredentialRepoImpl) GetByUserNameAndPassword(user model.UserCredential) error {
	isAuth := false
	for _, u := range *u.userCredDb {
		userName := user.GetUserName()
		userPassword := user.GetUserPassword()
		if userName == u.GetUserName() && userPassword == u.GetUserPassword() {
			isAuth = true
			break
		}
	}
	if isAuth {
		return nil
	}
	return utils.DataNotFoundError()
}

func NewUserCredentialRepo(userCredDb *[]model.UserCredential) UserCredentialRepo {
	return &userCredentialRepoImpl{
		userCredDb,
	}
}
