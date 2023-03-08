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

func (s *UsersService) SignUp(user models.User) (err error) {
	user.Password, err = s.hashing.Hash(user.Password)
	if err != nil {
		return fmt.Errorf("failed to create password hash, %v", err)
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		return fmt.Errorf("failed to create seller, %v", err)
	}

	return err
}

func (s *UsersService) SignIn(user models.User) (string, string, error) {
	passwordHash, err := s.hashing.Hash(user.Password)
	if err != nil {
		return "", "", err
	}

	userID, err := s.repo.GetByCredentials(user.Email, passwordHash)
	if err != nil {
		return "", "", fmt.Errorf("failed to find user with given credentials, %v", err)
	}

	accessToken, err := s.tokenManager.NewJWT(userID, s.accessTokenTTL)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.tokenManager.NewRefreshToken()
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
