package service

import (
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/models"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/repository"
)

type SellersService struct {
	repo repository.Sellers
}

func NewSellersService(repo repository.Sellers) *SellersService {
	return &SellersService{
		repo: repo,
	}
}

func (s *SellersService) Create(seller models.Seller) error {
	return s.repo.Create(seller)
}
