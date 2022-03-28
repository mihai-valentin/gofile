package service

import (
	"gofile/pkg/entity"
	"gofile/pkg/repository"
)

type PublicFile interface {
	UploadFiles(files entity.PublicFilesList) ([]*entity.FileEntity, error)
	GetFile(uuid string) (entity.PublicFile, error)
	DeleteFile(uuid string) error
}

type PrivateFile interface {
	UploadFiles(files entity.PrivateFilesList) ([]*entity.File, error)
	GetFile(uuid string, fileOwner entity.FileOwner) (entity.PrivateFile, error)
	DeleteFile(uuid string, fileOwner entity.FileOwner) error
}

type Service struct {
	PublicFile
	PrivateFile
}

func NewService(repositories *repository.Repository) *Service {
	return &Service{
		PublicFile:  NewPublicFileService(repositories.PublicFile),
		PrivateFile: NewPrivateFileService(repositories.PrivateFile),
	}
}
