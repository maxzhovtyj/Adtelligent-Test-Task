package repository

import (
	"database/sql"
	"fmt"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/models"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) CreateSeller(seller models.Seller) error {
	queryInsertSeller := fmt.Sprintf("INSERT INTO %s (name, phone_number) VALUES (?, ?)", sellersTable)

	row := r.db.QueryRow(queryInsertSeller, seller.Name, seller.PhoneNumber)
	if row.Err() != nil {
		return row.Err()
	}

	return nil
}
