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
	CreateUser(seller models.User) error
	GetByCredentials(email, password string) (int, error)
}

type Repository struct {
	Users
}

func New(db *sql.DB) *Repository {
	return &Repository{
		Users: NewUsersRepo(db),
	}
}
