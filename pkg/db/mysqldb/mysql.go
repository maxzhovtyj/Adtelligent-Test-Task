package mysqldb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const mySqlDriver = "mysql"

func NewClient(user, password, dbName string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@/%s", user, password, dbName)

	db, err := sql.Open(mySqlDriver, dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
