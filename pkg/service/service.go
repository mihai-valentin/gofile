package service

import (
	"gofile/pkg/entity"
	"gofile/pkg/handler"
)

type FileRepositoryInterface interface {
	CreateFile(file *entity.File) error
	FindFileByUuid(uuid string) (*entity.File, error)
	DeleteFileByUuid(uuid string) error
}

func New(repository FileRepositoryInterface, access string) *handler.FileServiceInterface {
	if access == "public" {
		return new(PublicFileManager)
	}

	return new(PrivateFileManager)
}
