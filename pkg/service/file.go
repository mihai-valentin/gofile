package service

import (
	"errors"
	"gofile/pkg/data"
	"gofile/pkg/entity"
	"os"
)

type FileManager struct {
	fileWriter   *FileWriter
	presetWriter *PresetWriter
	repository   FileRepositoryInterface
	storageRoot  string
}

func NewPublicFileManager(repository FileRepositoryInterface, storageRoot string) *FileManager {
	return &FileManager{
		fileWriter:   new(FileWriter),
		presetWriter: new(PresetWriter),
		repository:   repository,
		storageRoot:  storageRoot,
	}
}

func (s *FileManager) UploadFiles(filesUploadData []*data.UploadFileData) ([]*entity.File, error) {
	var files []*entity.File
	for _, fileUploadData := range filesUploadData {
		if err := s.uploadFile(fileUploadData); err != nil {
			return nil, err
		}

		file, err := s.repository.CreateFile(fileUploadData)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, nil
}

func (s *FileManager) uploadFile(fileUploadData *data.UploadFileData) error {
	filePath := s.buildFullPath(fileUploadData.GetPath())

	if fileUploadData.IsResizable() && fileUploadData.IsImage() {
		return s.presetWriter.writePreset(fileUploadData, filePath)
	}

	return s.fileWriter.writeFile(fileUploadData, filePath)
}

func (s *FileManager) GetFile(uuid string, ownerSign string) (string, error) {
	file, err := s.repository.FindByUuid(uuid)
	if err != nil {
		return "", err
	}

	if !s.canAccess(ownerSign, file) {
		return "", errors.New("no access")
	}

	return s.buildFullPath(file.Path), nil
}

func (s *FileManager) DeleteFile(uuid string, ownerSign string) error {
	filePath, err := s.GetFile(uuid, ownerSign)
	if err != nil {
		return err
	}

	if err := os.Remove(filePath); err != nil {
		return err
	}

	if err := s.repository.DeleteByUuid(uuid); err != nil {
		return err
	}

	return nil
}

func (s *FileManager) buildFullPath(path string) string {
	return s.storageRoot + "/" + path
}

func (s *FileManager) canAccess(ownerSign string, file *entity.File) bool {
	return file.Disk != "private" || file.OwnerSign == ownerSign
}
