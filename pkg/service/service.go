package service

import (
	"gofile/pkg/entity"
	"gofile/pkg/infrastructure"
	"gofile/pkg/repository"
)

type Error interface {
	error
	GetCode() int
}

type FileService interface {
	UploadFiles(filesUploadForm entity.FilesUploadForm) ([]*entity.File, error)
	GetFile(uuid string) (*entity.File, Error)
	DeleteFile(uuid string) Error
	MatchOwner(f *entity.File, o *entity.FileOwner) bool
}

type Service struct {
	PublicFileManager  FileService
	PrivateFileManager FileService
}

func NewService(repositories *repository.Repository, c *infrastructure.Config) *Service {
	return &Service{
		PublicFileManager:  NewFileService(repositories.FileRepositoryInterface, c.Get("access.public_mode")),
		PrivateFileManager: NewFileService(repositories.FileRepositoryInterface, c.Get("access.private_mode")),
	}
}
