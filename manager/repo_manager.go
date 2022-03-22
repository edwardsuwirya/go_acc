package manager

import (
	"enigmacamp.com/goacc/repository"
)

type RepoManager interface {
	ProductRepo() repository.ProductRepo
	UserCredentialRepo() repository.UserCredentialRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) ProductRepo() repository.ProductRepo {
	return repository.NewProductRepo(r.infra.SqlDb())
}

func (r *repoManager) UserCredentialRepo() repository.UserCredentialRepo {
	return repository.NewUserCredentialRepo(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
