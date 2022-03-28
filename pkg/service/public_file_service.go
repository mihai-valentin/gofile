package service

import (
	"errors"
	"fmt"
	"github.com/nfnt/resize"
	"gofile/pkg/entity"
	"gofile/pkg/repository"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

type ImageProcessor struct {
}

func (s *ImageProcessor) IsImage(f *entity.File) bool {
	return strings.HasPrefix(f.Source.Header.Get("Content-Type"), "image")
}

func (s *ImageProcessor) WritePresets(f *entity.File, presets []uint) ([]*entity.FileEntity, error) {
	if len(presets) == 0 {
		return nil, nil
	}

	source, err := os.Open(f.Entity.Path)
	if err != nil {
		return nil, err
	}

	original, _, err := image.Decode(source)
	if err != nil {
		return nil, err
	}
	source.Close()

	var uploaded []*entity.FileEntity

	for _, preset := range presets {
		thumbnail := resize.Resize(preset, 0, original, resize.Lanczos3)
		thumbnailName := fmt.Sprintf("%d_%s", preset, f.Source.Filename)
		thumbnailPath := fmt.Sprintf("storage/%s/%s_%s", f.Entity.Access, f.Entity.Uuid, thumbnailName)
		thumbnailFile, err := os.Create(thumbnailPath)

		if err != nil {
			return nil, err
		}

		switch strings.Split(f.Source.Header.Get("Content-Type"), "/")[1] {
		case "jpeg":
			err = jpeg.Encode(thumbnailFile, thumbnail, nil)
		case "png":
			err = png.Encode(thumbnailFile, thumbnail)
		case "gif":
			err = gif.Encode(thumbnailFile, thumbnail, nil)
		default:
			err = errors.New("image type is not supported yet")
		}

		if err != nil {
			return nil, err
		}

		thumbnailFile.Close()

		fileEntity := entity.NewPublicFileEntity(thumbnailName)
		fileEntity.Path = thumbnailPath

		uploaded = append(uploaded, fileEntity)
	}

	return uploaded, nil
}

type PublicFileService struct {
	*FileService
	*ImageProcessor
	repository repository.PublicFile
}

func NewPublicFileService(repository repository.PublicFile) *PublicFileService {
	return &PublicFileService{
		FileService:    new(FileService),
		ImageProcessor: new(ImageProcessor),
		repository:     repository,
	}
}

func (s *PublicFileService) UploadFiles(filesList entity.PublicFilesList) ([]*entity.FileEntity, error) {
	var files []*entity.File

	for _, formFile := range filesList.FormFiles {
		files = append(files, entity.NewPublicFile(formFile))
	}

	var uploaded []*entity.FileEntity

	for _, file := range files {
		s.repository.Create(file.Entity)

		if err := s.write(file); err != nil {
			return nil, err
		}

		uploaded = append(uploaded, file.Entity)

		if s.IsImage(file) && len(filesList.Presets) > 0 {
			uploadedPresets, err := s.WritePresets(file, filesList.Presets)
			if err != nil {
				return nil, err
			}

			for _, preset := range uploadedPresets {
				s.repository.Create(preset)
				uploaded = append(uploaded, preset)
			}
		}
	}

	return uploaded, nil
}

func (s *PublicFileService) GetFile(uuid string) (entity.PublicFile, error) {
	return s.repository.FindByUuid(uuid)
}

func (s *PublicFileService) DeleteFile(uuid string) error {
	file, err := s.repository.FindByUuid(uuid)
	if err != nil {
		return err
	}

	err = s.repository.DeleteByUuid(uuid)
	if err != nil {
		return err
	}

	return s.FileService.remove(file.Path)
}
