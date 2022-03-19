package manager

import "enigmacamp.com/goacc/repository"

type RepoManager interface {
	ProductRepo() repository.ProductRepo
	UserCredentialRepo() repository.UserCredentialRepo
}

type repoManager struct {
}

func (r *repoManager) ProductRepo() repository.ProductRepo {
	return repository.NewProductRepo()
}

func (r *repoManager) UserCredentialRepo() repository.UserCredentialRepo {
	return repository.NewUserCredentialRepo()
}

func NewRepoManager() RepoManager {
	return new(repoManager)
}
