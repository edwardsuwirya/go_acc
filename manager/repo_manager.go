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
	productDb []model.Product
	userDb    []model.UserCredential
}

func (r *repoManager) ProductRepo() repository.ProductRepo {
	return repository.NewProductRepo(&r.productDb)
}

func (r *repoManager) UserCredentialRepo() repository.UserCredentialRepo {
	return repository.NewUserCredentialRepo(&r.userDb)
}

func NewRepoManager() RepoManager {
	productDb := make([]model.Product, 0)
	userDb := []model.UserCredential{
		model.NewUserCredential("admin", "123"),
	}
	return &repoManager{
		productDb: productDb,
		userDb:    userDb,
	}
}
