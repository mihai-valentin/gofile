package service

import (
	"gofile/pkg/entity"
	"gofile/pkg/repository"
	"gofile/pkg/request"
)

type FileManager struct {
	*FileBuilder
	*StorageService
	*ImageProcessor
	access     string
	repository repository.FileRepositoryInterface
}

func NewFileService(repository repository.FileRepositoryInterface, mode string) *FileManager {
	return &FileManager{
		access:         mode,
		FileBuilder:    new(FileBuilder),
		StorageService: new(StorageService),
		ImageProcessor: new(ImageProcessor),
		repository:     repository,
	}
}

func (s *FileManager) UploadFiles(pubicFilesUploadRequest request.PubicFilesUploadRequest) ([]*entity.File, error) {
	var files []*entity.File
	presetsCount := len(pubicFilesUploadRequest.Presets)

	for _, formFile := range pubicFilesUploadRequest.FormFiles {
		formFileContentType := formFile.Header.Get("Content-Type")
		file := s.buildFile(s.access, formFile.Filename, formFileContentType)

		if err := s.repository.Create(file); err != nil {
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
			if err := s.repository.Create(preset); err != nil {
				return nil, err
			}

			files = append(files, preset)
		}
	}

	return files, nil
}

func (s *FileManager) GetFile(uuid string) (*entity.File, error) {
	file, err := s.repository.FindByUuidAndAccess(uuid, s.access)
	if file == nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *FileManager) DeleteFile(uuid string) error {
	file, err := s.GetFile(uuid)
	if err != nil {
		return err
	}

	if err := s.repository.DeleteByUuid(uuid); err != nil {
		return err
	}

	if err := s.remove(file.Path); err != nil {
		return err
	}

	return nil
}

func (s *FileManager) MatchOwner(f *entity.File, o *entity.FileOwner) bool {
	return o.Guid == f.OwnerGuid.String && o.Type == f.OwnerType.String
}
