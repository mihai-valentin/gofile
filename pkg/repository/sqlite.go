package repository

import (
	"github.com/jmoiron/sqlx"
)

const (
	filesTable = "files"
)

type Config struct {
	File string
}

func NewSqliteDB(c Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", c.File)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
