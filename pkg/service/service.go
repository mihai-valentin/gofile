package service

import (
	"gofile/pkg/entity"
	"gofile/pkg/repository"
)

type FileRepositoryInterface interface {
	CreateFile(fileData repository.FileDataInterface) (*entity.File, error)
	FindByUuid(uuid string) (*entity.File, error)
	DeleteByUuid(uuid string) error
}

type Service struct {
	*FileManager
}

func New(repository FileRepositoryInterface, storageRoot string) *Service {
	return &Service{
		FileManager: NewPublicFileManager(repository, storageRoot),
	}
}
