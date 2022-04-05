package service

import (
	"errors"
	"gofile/pkg/contracts"
	"os"
)

type FileManager struct {
	fileWriter     *FileWriter
	presetWriter   *PresetWriter
	storageManager *StorageManager
	repository     contracts.FileRepositoryInterface
}

func NewPublicFileManager(repository contracts.FileRepositoryInterface, storageRoot string) *FileManager {
	return &FileManager{
		fileWriter:     new(FileWriter),
		presetWriter:   new(PresetWriter),
		storageManager: NewStorageManager(storageRoot),
		repository:     repository,
	}
}

func (s *FileManager) UploadFiles(filesUploadData []contracts.FileUploadDataInterface) (map[string]string, error) {
	filesUrls := map[string]string{}
	for _, fileUploadData := range filesUploadData {
		if err := s.uploadFile(fileUploadData); err != nil {
			return nil, err
		}

		file, err := s.repository.StoreFile(fileUploadData)
		if err != nil {
			return nil, err
		}
		filesUrls[file.GetName()] = s.storageManager.buildFileAccessUrl(file.GetUuid())
	}

	return filesUrls, nil
}

func (s *FileManager) uploadFile(fileUploadData contracts.FileUploadDataInterface) error {
	filePath := s.storageManager.buildFileFullPath(fileUploadData.GetPath())

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

	if !s.storageManager.checkAccessByOwnSign(ownerSign, file) {
		return "", errors.New("no access")
	}

	return s.storageManager.buildFileFullPath(file.GetPath()), nil
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
