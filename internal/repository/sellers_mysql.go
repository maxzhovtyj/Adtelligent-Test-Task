package repository

import (
	"database/sql"
	"fmt"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/models"
)

type SellersRepo struct {
	db *sql.DB
}

func NewSellersRepo(db *sql.DB) *SellersRepo {
	return &SellersRepo{
		db: db,
	}
}

func (r *SellersRepo) Create(seller models.Seller) error {
	queryInsertSeller := fmt.Sprintf("INSERT INTO %s (name, phone_number) VALUES (?, ?)", sellersTable)

	row := r.db.QueryRow(queryInsertSeller, seller.Name, seller.PhoneNumber)
	if row.Err() != nil {
		return row.Err()
	}

	return nil
}
