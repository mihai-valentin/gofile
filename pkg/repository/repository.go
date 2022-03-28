package repository

import (
	"gofile/pkg/entity"
	"gofile/pkg/infrastructure"
)

type PublicFile interface {
	Create(file *entity.FileEntity)
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

func NewRepository(d *infrastructure.Database) *Repository {
	return &Repository{
		PublicFile:  NewPublicFileRepository(d),
		PrivateFile: NewPrivateFileRepository(d),
	}
}
