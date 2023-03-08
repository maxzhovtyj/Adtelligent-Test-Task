package service

import (
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/models"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/repository"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/pkg/auth"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/pkg/hash"
	"time"
)

type Users interface {
	SignUp(seller models.Seller) error
	SignIn(seller models.Seller) error
}

type Service struct {
	Users
}

func New(
	repo *repository.Repository,
	manager auth.TokenManager,
	accessTokenTTL, refreshTokenTTL time.Duration,
	hashing hash.PasswordHashing,
) *Service {
	return &Service{
		Users: NewUsersService(repo.Users, manager, accessTokenTTL, refreshTokenTTL, hashing),
	}
}
