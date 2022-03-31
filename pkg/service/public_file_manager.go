package service

import (
	"gofile/pkg/entity"
	"gofile/pkg/repository"
	"gofile/pkg/request"
)

type PublicFileManager struct {
	*FileBuilder
	*StorageService
	*ImageProcessor
	access     string
	repository repository.PublicFileSqlite
}

func NewPublicFileManager(repository repository.PublicFileSqlite) *PublicFileManager {
	return &PublicFileManager{
		FileBuilder:    new(FileBuilder),
		StorageService: new(StorageService),
		ImageProcessor: new(ImageProcessor),
		repository:     repository,
	}
}

func (s *PublicFileManager) UploadPublicFiles(pubicFilesUploadRequest request.PubicFilesUploadRequest) ([]*entity.File, error) {
	var files []*entity.File
	presetsCount := len(pubicFilesUploadRequest.Presets)

	for _, formFile := range pubicFilesUploadRequest.FormFiles {
		formFileContentType := formFile.Header.Get("Content-Type")
		file := s.buildFile(s.access, formFile.Filename, formFileContentType)

		if err := s.repository.CreatePublicFile(file); err != nil {
			return nil, err
		}

		if err := s.write(formFile, file.Path); err != nil {
			return nil, err
		}

		files = append(files, file)

		if presetsCount <= 0 || !file.IsAnImage() {
			continue
		}

		presets, err := s.WritePresets(file, pubicFilesUploadRequest.Presets)
		if err != nil {
			return nil, err
		}

		for _, preset := range presets {
			if err := s.repository.CreatePublicFile(preset); err != nil {
				return nil, err
			}

			files = append(files, preset)
		}
	}

	return files, nil
}

func (s *PublicFileManager) GetPublicFile(uuid string) (*entity.File, error) {
	file, err := s.repository.FindPublicFileByUuid(uuid)
	if file == nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *PublicFileManager) DeletePublicFile(uuid string) error {
	file, err := s.GetPublicFile(uuid)
	if err != nil {
		return err
	}

	if err := s.repository.DeletePublicFileByUuid(uuid); err != nil {
		return err
	}

	if err := s.remove(file.Path); err != nil {
		return err
	}

	return nil
}
