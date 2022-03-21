package usecase

import (
	"enigmacamp.com/goacc/model"
	"enigmacamp.com/goacc/repository"
)

type ProductRegistrationUseCase interface {
	Register(productCode, productName string) model.Product
}

type productRegistrationUseCase struct {
	repo repository.ProductRepo
}

func (a *productRegistrationUseCase) Register(productCode, productName string) model.Product {
	newProduct := model.NewProduct(productCode, productName)
	a.repo.Insert(newProduct)
	return newProduct
}

func NewProductRegistrationUseCase(repo repository.ProductRepo) ProductRegistrationUseCase {
	return &productRegistrationUseCase{
		repo: repo,
	}
}
