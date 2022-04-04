package repository

import (
	"gofile/internal"
)

type Repository struct {
	*FileSqlite
}

func New(db *internal.Database) *Repository {
	return &Repository{
		NewPublicFileSqlite(db.DB),
	}
}
