package usecase

import (
	"enigmacamp.com/goacc/model"
	"enigmacamp.com/goacc/repository"
)

type AuthenticationUseCase interface {
	Login(userName string, userPassword string) bool
}

type authenticationUseCase struct {
	repo repository.UserCredentialRepo
}

func (a *authenticationUseCase) Login(userName string, userPassword string) bool {
	userAuth := model.NewUserCredential(userName, userPassword)
	return a.repo.GetByUserNameAndPassword(userAuth)
}

func NewAuthenticationUseCase(repo repository.UserCredentialRepo) AuthenticationUseCase {
	return &authenticationUseCase{
		repo: repo,
	}
}
