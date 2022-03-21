package repository

import (
	"enigmacamp.com/goacc/model"
)

type userCredentialRepoImpl struct {
	userCredDb *[]model.UserCredential
}

func (u *userCredentialRepoImpl) GetByUserNameAndPassword(user model.UserCredential) (isAuth bool) {
	isAuth = false
	for _, u := range *u.userCredDb {
		userName := user.GetUserName()
		userPassword := user.GetUserPassword()
		if userName == u.GetUserName() && userPassword == u.GetUserPassword() {
			isAuth = true
			break
		}
	}
	return
}

func NewUserCredentialRepo(userCredDb *[]model.UserCredential) UserCredentialRepo {
	return &userCredentialRepoImpl{
		userCredDb,
	}
}
