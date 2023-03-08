package repository

import (
	"database/sql"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/models"
)

const (
	ordersTable    = "orders"
	productsTable  = "products"
	customersTable = "customers"
	sellersTable   = "sellers"
)

type Users interface {
	CreateSeller(seller models.Seller) error
}

type Repository struct {
	Users
}

func New(db *sql.DB) *Repository {
	return &Repository{
		Users: NewUsersRepo(db),
	}
}
