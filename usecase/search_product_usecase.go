package usecase

import (
	"enigmacamp.com/goacc/model"
	"enigmacamp.com/goacc/repository"
)

type SearchProductUseCase interface {
	Search(productCode string) []model.Product
}

type searchProductUseCase struct {
	repo repository.ProductRepo
}

func (a *searchProductUseCase) Search(productCode string) []model.Product {
	if len(productCode) == 0 {
		return a.repo.GetAll()
	}
	result := a.repo.GetByProductCode(productCode)
	if len(result.ProductCode) == 0 {
		return nil
	} else {
		return []model.Product{result}
	}
}

func NewSearchProductUseCase(repo repository.ProductRepo) SearchProductUseCase {
	return &searchProductUseCase{
		repo: repo,
	}
}
