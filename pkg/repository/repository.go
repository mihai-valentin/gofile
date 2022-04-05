package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	*FileSqlite
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		NewPublicFileSqlite(db),
	}
}
