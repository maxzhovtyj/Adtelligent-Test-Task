package service

import (
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

func (s *UsersService) SignUp(seller models.Seller) error {

	return nil
}

func (s *UsersService) SignIn(seller models.Seller) error {
	return nil
}
