package repository

import "gofile/pkg/infrastructure"

type Repository struct {
	*PublicFileSqlite
	*PrivateFileSqlite
}

func New(db *infrastructure.Database) *Repository {
	return &Repository{
		NewPublicFileSqlite(db.DB),
		NewPrivateFileSqlite(db.DB),
	}
}
