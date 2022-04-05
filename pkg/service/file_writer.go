package service

import (
	"gofile/pkg/contracts"
	"io"
	"os"
)

type FileWriter struct {
}

func (w *FileWriter) writeFile(fileData contracts.FileUploadDataInterface, path string) error {
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
