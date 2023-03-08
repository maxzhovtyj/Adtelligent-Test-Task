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

type repository struct {
	Users
}

func New(db *sql.DB) *repository {
	return &repository{
		Users: NewUsersRepo(db),
	}
}
