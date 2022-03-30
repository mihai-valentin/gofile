package service

import (
	"errors"
	"github.com/nfnt/resize"
	"gofile/pkg/entity"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

type ImageProcessor struct {
	*FileBuilder
}

func (processor *ImageProcessor) WritePresets(f *entity.File, presets []uint) ([]*entity.File, error) {
	if len(presets) == 0 {
		return nil, nil
	}

	file, err := decodeImage(f.Path)
	if err != nil {
		return nil, err
	}

	var uploaded []*entity.File
	for _, preset := range presets {
		presetFile := processor.buildPresetFile(f, preset)

		if err := resizeAndSaveFile(file, presetFile, preset); err != nil {
			return nil, err
		}

		uploaded = append(uploaded, presetFile)
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

func resizeAndSaveFile(file image.Image, presetFile *entity.File, scale uint) error {
	preset := resize.Resize(scale, 0, file, resize.Lanczos3)

	dist, err := os.Create(presetFile.Path)
	if err != nil {
		return err
	}

	if err := encodeImage(presetFile.Encoding, preset, dist); err != nil {
		return err
	}

	return dist.Close()
}
