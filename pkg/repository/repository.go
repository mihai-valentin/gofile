package repository

import "gofile/pkg/infrastructure"

type Repository struct {
	*FileSqlite
}

func New(db *infrastructure.Database) *Repository {
	return &Repository{
		NewPublicFileSqlite(db.DB),
	}
}
