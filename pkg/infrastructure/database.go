package infrastructure

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

type Database struct {
	DB     *sqlx.DB
	Driver string
	Schema string
}

func NewSqliteDB() (*Database, error) {
	driver := os.Getenv("DB_DRIVER")
	schema := os.Getenv("DB_SCHEMA")

	db, err := sqlx.Open(driver, schema)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Database{
		DB:     db,
		Driver: driver,
		Schema: schema,
	}, nil
}

func (d *Database) Close() error {
	return d.DB.Close()
}
