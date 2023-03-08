package repository

import "database/sql"

type Users interface {
	Create()
}

type repository struct {
	Users
}

func New(db *sql.DB) *repository {
	return &repository{
		Users: NewUsersRepo(db),
	}
}
