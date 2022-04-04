package service

import (
	"io"
	"mime/multipart"
	"os"
)

type FileDataInterface interface {
	GetSource() *multipart.FileHeader
}

type FileWriter struct {
}

func (w *FileWriter) writeFile(fileData FileDataInterface, path string) error {
	src, err := fileData.GetSource().Open()
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
