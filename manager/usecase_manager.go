package manager

import "enigmacamp.com/goacc/usecase"

type UseCaseManager interface {
	UserLoginUseCase() usecase.AuthenticationUseCase
	ProductRegistrationUseCase() usecase.ProductRegistrationUseCase
	SearchProductUseCase() usecase.SearchProductUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) UserLoginUseCase() usecase.AuthenticationUseCase {
	return usecase.NewAuthenticationUseCase(u.repo.UserCredentialRepo())
}

func (u *useCaseManager) ProductRegistrationUseCase() usecase.ProductRegistrationUseCase {
	return usecase.NewProductRegistrationUseCase(u.repo.ProductRepo())
}

func (u *useCaseManager) SearchProductUseCase() usecase.SearchProductUseCase {
	return usecase.NewSearchProductUseCase(u.repo.ProductRepo())
}
func NewUseCaseManager(manager RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: manager,
	}
}
