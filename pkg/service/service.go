package service

import "gofile/pkg/contracts"

type Service struct {
	*FileManager
}

func New(repository contracts.FileRepositoryInterface, storageRoot string) *Service {
	return &Service{
		FileManager: NewPublicFileManager(repository, storageRoot),
	}
}
