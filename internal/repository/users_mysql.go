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

func (r *UsersRepo) CreateUser(user models.User) error {
	queryInsertSeller := fmt.Sprintf(
		"INSERT INTO %s (email, password) VALUES (?, ?)",
		usersTable,
	)

	row := r.db.QueryRow(queryInsertSeller, user.Email, user.Password)
	if row.Err() != nil {
		return row.Err()
	}

	return nil
}

func (r *UsersRepo) GetByCredentials(email, password string) (int, error) {
	var userID int

	queryGetUser := fmt.Sprintf("SELECT id FROM users WHERE email = ? AND password = ?")

	row := r.db.QueryRow(queryGetUser, email, password)
	if row.Err() != nil {
		return 0, row.Err()
	}

	if err := row.Scan(&userID); err != nil {
		return 0, err
	}

	return userID, nil
}
