package service

import (
	"fmt"
	"gofile/pkg/contracts"
	"os"
)

const privateDisk = "private"

type StorageManager struct {
	storageRoot string
}

func NewStorageManager(storageRoot string) *StorageManager {
	return &StorageManager{storageRoot}
}

func (s *StorageManager) buildFileFullPath(path string) string {
	return s.storageRoot + "/" + path
}

func (s *StorageManager) buildFileAccessUrl(uuid string) string {
	return fmt.Sprintf("%s/api/files/%s", os.Getenv("SERVICE_HOST"), uuid)
}

func (s *StorageManager) checkAccessByOwnSign(ownerSign string, file contracts.FileEntityInterface) bool {
	return file.GetDisk() != privateDisk || file.GetOwnerSign() == ownerSign
}
