package repository

import (
	"github.com/jmoiron/sqlx"
	"gofile/pkg/entity"
)

type PublicFile interface {
	Create(file entity.File)
	DeleteByUuid(uuid string) error
	FindByUuid(uuid string) (entity.PublicFile, error)
}

type PrivateFile interface {
	Create(file entity.File)
	DeleteByUuid(uuid string) error
	FindByUuid(uuid string) (entity.PrivateFile, error)
}

type Repository struct {
	PublicFile
	PrivateFile
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		PublicFile:  NewPublicFileRepository(db),
		PrivateFile: NewPrivateFileRepository(db),
	}
}
