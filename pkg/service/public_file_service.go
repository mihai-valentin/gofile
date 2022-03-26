package service

import (
	"gofile/pkg/entity"
	"gofile/pkg/repository"
)

type PublicFileService struct {
	*FileService
	repository repository.PublicFile
}

func (s *PublicFileService) UploadFiles(files entity.PublicFilesList) ([]*entity.File, error) {
	var uploadedFiles []*entity.File

	for _, file := range files.Payload {
		publicFile := entity.NewPublicFile(file)
		s.repository.Create(*publicFile)

		err := new(FileService).write(publicFile)
		if err != nil {
			return nil, err
		}

		uploadedFiles = append(uploadedFiles, publicFile)
	}

	return uploadedFiles, nil
}

func (s *PublicFileService) GetFile(uuid string) (entity.PublicFile, error) {
	return s.repository.FindByUuid(uuid)
}

func (s *PublicFileService) DeleteFile(uuid string) error {
	file, err := s.repository.FindByUuid(uuid)
	if err != nil {
		return err
	}

	err = s.repository.DeleteByUuid(uuid)
	if err != nil {
		return err
	}

	return s.FileService.remove(file.Path)
}

func NewPublicFileService(repository repository.PublicFile) *PublicFileService {
	return &PublicFileService{repository: repository}
}
