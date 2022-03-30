package repository

import (
	"gofile/pkg/entity"
	"gofile/pkg/infrastructure"
)

type FileRepositoryInterface interface {
	Create(file *entity.File) error
	DeleteByUuid(uuid string) error
	FindByUuidAndAccess(uuid string, access string) (*entity.File, error)
}

type Repository struct {
	FileRepositoryInterface
}

func NewRepository(db *infrastructure.Database) *Repository {
	return &Repository{
		FileRepositoryInterface: NewFileRepository(db),
	}
}
