package manager

import (
	"enigmacamp.com/goacc/model"
	"enigmacamp.com/goacc/repository"
)

type RepoManager interface {
	ProductRepo() repository.ProductRepo
	UserCredentialRepo() repository.UserCredentialRepo
}

type repoManager struct {
}

func (r *repoManager) ProductRepo() repository.ProductRepo {
	productDb := make([]model.Product, 0)
	return repository.NewProductRepo(&productDb)
}

func (r *repoManager) UserCredentialRepo() repository.UserCredentialRepo {
	userDb := []model.UserCredential{
		model.NewUserCredential("admin", "123"),
	}
	return repository.NewUserCredentialRepo(&userDb)
}

func NewRepoManager() RepoManager {
	return new(repoManager)
}
