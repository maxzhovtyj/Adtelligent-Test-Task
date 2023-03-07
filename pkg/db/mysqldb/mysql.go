package mysqldb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func NewClient() (*sql.DB, error) {
	db, err := sql.Open("mysql", "mysql:30042003@/adtelligent-db")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, err
}
