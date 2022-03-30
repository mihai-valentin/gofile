package service

import (
	"database/sql"
	"github.com/rs/xid"
	"gofile/pkg/entity"
	"gofile/pkg/errors"
	"gofile/pkg/repository"
	"strings"
	"time"
)

type FileManager struct {
	mode string
	*StorageService
	*ImageProcessor
	repository repository.FileRepositoryInterface
}

func NewFileService(repository repository.FileRepositoryInterface, mode string) *FileManager {
	return &FileManager{
		mode:           mode,
		StorageService: new(StorageService),
		ImageProcessor: new(ImageProcessor),
		repository:     repository,
	}
}

func (s *FileManager) buildFile(filename string, contentType string) *entity.File {
	uuid := xid.New().String()
	path := "storage/" + s.mode + "/" + uuid + "_" + filename
	contentTypeParts := strings.Split(contentType, "/")

	return &entity.File{
		Uuid:        uuid,
		Name:        filename,
		Access:      s.mode,
		ContentType: contentType,
		Type:        contentTypeParts[0],
		Encoding:    contentTypeParts[1],
		Path:        path,
		CreatedAt:   time.Now(),
	}
}

func (s *FileManager) UploadFiles(filesUploadForm entity.FilesUploadForm) ([]*entity.File, error) {
	var files []*entity.File

	for _, formFile := range filesUploadForm.FormFiles {
		file := s.buildFile(formFile.Filename, formFile.Header.Get("Content-Type"))

		file.OwnerGuid = sql.NullString{String: filesUploadForm.OwnerGuid, Valid: true}
		file.OwnerType = sql.NullString{String: filesUploadForm.OwnerType, Valid: true}

		if err := s.repository.Create(file); err != nil {
			return nil, err
		}

		if err := s.write(formFile, file.Path); err != nil {
			return nil, err
		}

		if len(filesUploadForm.Presets) > 0 && file.Type == "image" {
			uploadedPresets, err := s.WritePresets(file, filesUploadForm.Presets)
			if err != nil {
				return nil, err
			}

			for _, preset := range uploadedPresets {
				if err := s.repository.Create(preset); err != nil {
					return nil, err
				}

				files = append(files, preset)
			}
		}

		files = append(files, file)
	}

	return files, nil
}

func (s *FileManager) GetFile(uuid string) (*entity.File, Error) {
	file, err := s.repository.FindByUuidAndAccess(uuid, s.mode)
	if file == nil {
		return nil, errors.NewNotFoundError(err)
	}

	if err != nil {
		return nil, errors.NewInternalServerError(err)
	}

	return file, nil
}

func (s *FileManager) DeleteFile(uuid string) Error {
	file, err := s.GetFile(uuid)
	if err != nil {
		return err
	}

	dbErr := s.repository.DeleteByUuid(uuid)
	if dbErr != nil {
		return errors.NewInternalServerError(dbErr)
	}

	osErr := s.remove(file.Path)
	if osErr != nil {
		return errors.NewInternalServerError(osErr)
	}

	return nil
}

func (s *FileManager) MatchOwner(f *entity.File, o *entity.FileOwner) bool {
	return o.Guid == f.OwnerGuid.String && o.Type == f.OwnerType.String
}
