package repository

import (
	"database/sql"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/models"
)

const (
	usersTable     = "users"
	sellersTable   = "sellers"
	customersTable = "customers"
	ordersTable    = "orders"
	productsTable  = "products"
)

type Users interface {
	CreateUser(user models.User) error
	GetByCredentials(email, password string) (int, error)
}

type Sellers interface {
	Create(seller models.Seller) error
}

type Products interface {
	Get(id int) (models.Product, error)
	Create(product models.Product) error
	Delete(productID int) error
	Update(product models.Product) error
}

type Repository struct {
	Users
	Sellers
	Products
}

func New(db *sql.DB) *Repository {
	return &Repository{
		Users:    NewUsersRepo(db),
		Sellers:  NewSellersRepo(db),
		Products: NewProductsRepo(db),
	}
}
