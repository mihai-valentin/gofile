package service

import (
	"errors"
	"gofile/pkg/entity"
	"gofile/pkg/repository"
)

type PrivateFileService struct {
	*FileService
	repository repository.PrivateFile
}

func (s *PrivateFileService) UploadFiles(files entity.PrivateFilesList) ([]*entity.File, error) {
	var uploadedFiles []*entity.File

	for _, file := range files.Payload {
		privateFile := entity.NewPrivateFile(file, files.FileOwner)
		s.repository.Create(*privateFile)

		err := s.FileService.write(privateFile)
		if err != nil {
			return nil, err
		}

		uploadedFiles = append(uploadedFiles, privateFile)
	}

	return uploadedFiles, nil
}

func (s *PrivateFileService) GetFile(uuid string, fileOwner entity.FileOwner) (entity.PrivateFile, error) {
	file, err := s.repository.FindByUuid(uuid)

	if err != nil {
		return entity.PrivateFile{}, err
	}

	if !s.matchOwner(file, fileOwner) {
		return entity.PrivateFile{}, errors.New("no access")
	}

	return file, nil
}

func (s *PrivateFileService) DeleteFile(uuid string, fileOwner entity.FileOwner) error {
	file, err := s.repository.FindByUuid(uuid)
	if err != nil {
		return err
	}

	if !s.matchOwner(file, fileOwner) {
		return errors.New("no access")
	}

	err = s.repository.DeleteByUuid(uuid)
	if err != nil {
		return err
	}

	return s.FileService.remove(file.Path)
}

func (s *PrivateFileService) matchOwner(f entity.PrivateFile, o entity.FileOwner) bool {
	return o.Uuid == f.OwnerUuid.String && o.Type == f.OwnerType.String
}

func NewPrivateFileService(repository repository.PrivateFile) *PrivateFileService {
	return &PrivateFileService{
		FileService: newFileService(),
		repository:  repository,
	}
}
