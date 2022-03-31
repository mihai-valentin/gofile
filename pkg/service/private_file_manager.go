package service

import (
	"gofile/pkg/entity"
	"gofile/pkg/repository"
	"gofile/pkg/request"
)

type PrivateFileManager struct {
	*FileBuilder
	*StorageService
	*ImageProcessor
	access     string
	repository repository.PrivateFileSqlite
}

func NewPrivateFileManager(repository repository.PrivateFileSqlite) *PrivateFileManager {
	return &PrivateFileManager{
		FileBuilder:    new(FileBuilder),
		StorageService: new(StorageService),
		ImageProcessor: new(ImageProcessor),
		repository:     repository,
	}
}

func (s *PrivateFileManager) UploadPrivateFiles(request request.PrivateFilesUploadRequest) ([]*entity.File, error) {
	var files []*entity.File
	presetsCount := len(request.Presets)

	for _, formFile := range request.FormFiles {
		formFileContentType := formFile.Header.Get("Content-Type")
		file := s.buildFile(s.access, formFile.Filename, formFileContentType)

		if err := s.repository.CreatePrivateFile(file); err != nil {
			return nil, err
		}

		if err := s.write(formFile, file.Path); err != nil {
			return nil, err
		}

		files = append(files, file)

		if presetsCount <= 0 || !file.IsAnImage() {
			continue
		}

		presets, err := s.WritePresets(file, request.Presets)
		if err != nil {
			return nil, err
		}

		for _, preset := range presets {
			if err := s.repository.CreatePrivateFile(preset); err != nil {
				return nil, err
			}

			files = append(files, preset)
		}
	}

	return files, nil
}

func (s *PrivateFileManager) GetPrivateFile(uuid string, request request.PrivateFileAccessRequest) (*entity.File, error) {
	file, err := s.repository.FindPrivateFileByUuidAndOwner(uuid, request.OwnerGuid, request.OwnerType)
	if file == nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *PrivateFileManager) DeletePrivateFile(uuid string, request request.PrivateFileAccessRequest) error {
	file, err := s.GetPrivateFile(uuid, request)
	if err != nil {
		return err
	}

	if err := s.repository.DeletePrivateFileByUuidAndOwner(uuid, request.OwnerGuid, request.OwnerType); err != nil {
		return err
	}

	if err := s.remove(file.Path); err != nil {
		return err
	}

	return nil
}
