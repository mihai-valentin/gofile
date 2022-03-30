package service

import (
	"errors"
	"fmt"
	"github.com/nfnt/resize"
	"github.com/rs/xid"
	"gofile/pkg/entity"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"time"
)

type ImageProcessor struct {
}

func buildThumbnail(original *entity.File, delta uint) *entity.File {
	filename := fmt.Sprintf("%d_%s", delta, original.Name)
	path := fmt.Sprintf("storage/%s/%s_%s", original.Access, original.Uuid, filename)
	uuid := xid.New().String()

	return &entity.File{
		Uuid:        uuid,
		Name:        filename,
		Access:      original.Access,
		ContentType: original.ContentType,
		Type:        original.Type,
		Encoding:    original.Encoding,
		Path:        path,
		OwnerGuid:   original.OwnerGuid,
		OwnerType:   original.OwnerType,
		CreatedAt:   time.Now(),
	}
}

func (processor *ImageProcessor) WritePresets(f *entity.File, presets []uint) ([]*entity.File, error) {
	if len(presets) == 0 {
		return nil, nil
	}

	original, err := decodeImage(f.Path)
	if err != nil {
		return nil, err
	}

	var uploaded []*entity.File

	for _, preset := range presets {
		thumbnail := buildThumbnail(f, preset)

		if err := writePreset(original, thumbnail, preset); err != nil {
			return nil, err
		}

		uploaded = append(uploaded, thumbnail)
	}

	return uploaded, nil
}

func decodeImage(path string) (image.Image, error) {
	source, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	decodedImage, _, err := image.Decode(source)
	if err != nil {
		return nil, err
	}

	err = source.Close()
	if err != nil {
		return nil, err
	}

	return decodedImage, nil
}

func encodeImage(encoding string, i image.Image, f *os.File) error {
	switch encoding {
	case "jpeg":
		return jpeg.Encode(f, i, nil)
	case "png":
		return png.Encode(f, i)
	case "gif":
		return gif.Encode(f, i, nil)
	default:
		return errors.New("image type is not supported yet")
	}
}

func writePreset(original image.Image, thumbnail *entity.File, delta uint) error {
	preset := resize.Resize(delta, 0, original, resize.Lanczos3)

	file, err := os.Create(thumbnail.Path)
	if err != nil {
		return err
	}

	if err := encodeImage(thumbnail.Encoding, preset, file); err != nil {
		return err
	}

	return file.Close()
}
