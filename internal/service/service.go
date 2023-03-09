package service

import (
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/models"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/repository"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/pkg/auth"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/pkg/hash"
	"time"
)

type Users interface {
	SignUp(user models.User) error
	SignIn(user models.User) (string, string, error)
}

type Sellers interface {
	Create(seller models.Seller) error
}

type Products interface {
	Get(id int) (models.Product, error)
	Create(product models.Product) error
	Update(product models.Product) error
	Delete(productID int) error
}

type Service struct {
	Users
	Sellers
	Products
}

func New(
	repo *repository.Repository,
	manager auth.TokenManager,
	accessTokenTTL, refreshTokenTTL time.Duration,
	hashing hash.PasswordHashing,
) *Service {
	return &Service{
		Users:    NewUsersService(repo.Users, manager, accessTokenTTL, refreshTokenTTL, hashing),
		Sellers:  NewSellersService(repo.Sellers),
		Products: NewProductsService(repo.Products),
	}
}
