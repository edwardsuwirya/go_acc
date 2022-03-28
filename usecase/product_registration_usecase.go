package usecase

import (
	"enigmacamp.com/goacc/model"
	"enigmacamp.com/goacc/repository"
)

type ProductRegistrationUseCase interface {
	Register(productCode, productName, categoryId string) (model.Product, error)
}

type productRegistrationUseCase struct {
	repo repository.ProductRepo
}

func (a *productRegistrationUseCase) Register(productCode, productName, categoryId string) (model.Product, error) {
	newProduct := model.NewProduct(productCode, productName, categoryId)
	return a.repo.Insert(newProduct)
}

func NewProductRegistrationUseCase(repo repository.ProductRepo) ProductRegistrationUseCase {
	return &productRegistrationUseCase{
		repo: repo,
	}
}
