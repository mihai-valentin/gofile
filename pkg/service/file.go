package service

import (
	gofile "gofile/pkg/entity"
	"io"
	"os"
)

type FileService struct {
}

func (s *FileService) write(f *gofile.File) error {
	src, err := f.Source.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(f.Entity.Path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func (s *FileService) remove(path string) error {
	return os.Remove(path)
}

func newFileService() *FileService {
	return &FileService{}
}
