package service

import (
	"io"
	"mime/multipart"
	"os"
)

type StorageService struct {
}

func (s *StorageService) write(f *multipart.FileHeader, path string) error {
	src, err := f.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func (s *StorageService) remove(path string) error {
	return os.Remove(path)
}
