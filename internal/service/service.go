package service

import (
	"github.com/maxzhovtyj/Adtelligent-Test-Task/pkg/auth"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/pkg/hash"
	"time"
)

type Service struct {
	tokenManager auth.TokenManager

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration

	hashing hash.PasswordHashing
}

func New() *Service {
	return &Service{}
}
