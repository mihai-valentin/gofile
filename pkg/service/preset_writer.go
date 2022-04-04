package service

import (
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"mime/multipart"
	"os"
)

type ImageFileDataInterface interface {
	GetSource() *multipart.FileHeader
	GetScale() uint
	GetEncoding() string
}

type PresetWriter struct {
}

func (w *PresetWriter) writePreset(fileData ImageFileDataInterface, path string) error {
	src, err := fileData.GetSource().Open()
	if err != nil {
		return err
	}
	defer src.Close()

	imageContent, _, err := image.Decode(src)
	if err != nil {
		return err
	}

	resizedImage := resize.Resize(fileData.GetScale(), 0, imageContent, resize.Lanczos3)

	out, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	return w.encodePreset(fileData.GetEncoding(), out, resizedImage)
}

func (w *PresetWriter) encodePreset(encoding string, dist *os.File, src image.Image) error {
	switch encoding {
	case "jpeg":
		return jpeg.Encode(dist, src, nil)
	case "png":
		return png.Encode(dist, src)
	case "gif":
		return gif.Encode(dist, src, nil)
	default:
		return errors.New("image type is not supported yet")
	}
}
