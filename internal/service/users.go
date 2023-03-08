package service

import (
	"fmt"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/models"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/repository"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/pkg/auth"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/pkg/hash"
	"time"
)

type UsersService struct {
	repo         repository.Users
	tokenManager auth.TokenManager

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration

	hashing hash.PasswordHashing
}

func NewUsersService(
	repo repository.Users,
	manager auth.TokenManager,
	accessTokenTTL, refreshTokenTTL time.Duration,
	hashing hash.PasswordHashing,
) *UsersService {
	return &UsersService{
		repo:            repo,
		tokenManager:    manager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
		hashing:         hashing,
	}
}

func (s *UsersService) SignUpSeller(seller models.Seller) (err error) {
	seller.Password, err = s.hashing.Hash(seller.Password)
	if err != nil {
		return fmt.Errorf("failed to create password hash, %v", err)
	}

	err = s.repo.CreateSeller(seller)
	if err != nil {
		return fmt.Errorf("failed to create seller, %v", err)
	}

	return err
}

func (s *UsersService) SignInSeller(seller models.Seller) error {
	return nil
}

func (s *UsersService) SignUpCustomer(customer models.Customer) error {
	return nil
}

func (s *UsersService) SignInCustomer(customer models.Customer) error {
	return nil
}
