package service

import (
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/models"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/repository"
)

type ProductsService struct {
	repo repository.Products
}

func NewProductsService(repo repository.Products) *ProductsService {
	return &ProductsService{repo: repo}
}

func (s *ProductsService) Create(product models.Product) error {
	return s.repo.Create(product)
}

func (s *ProductsService) Delete(productID int) error {
	return s.repo.Delete(productID)
}
