package mapper

import (
	"fmt"
	"gofile/pkg/data"
	"mime/multipart"
	"strings"
)

type UploadRequestInterface interface {
	GetFormFiles() []*multipart.FileHeader
	GetPresets() []uint
	GetDisk() string
	GetOwnerSign() string
}

type UploadDataMapper struct {
}

func (m *UploadDataMapper) MapFromRequest(r UploadRequestInterface) []*data.UploadFileData {
	formFiles := r.GetFormFiles()
	arePresetsPresent := len(r.GetPresets()) > 0

	var images []*multipart.FileHeader
	var filesData []*data.UploadFileData

	for _, formFile := range formFiles {
		uploadFileData := data.NewUploadFileData(formFile, formFile.Filename, r.GetDisk(), r.GetOwnerSign(), 0)
		filesData = append(filesData, uploadFileData)

		if !arePresetsPresent {
			continue
		}

		formFileContentType := formFile.Header.Get("Content-Type")
		if strings.HasPrefix(formFileContentType, "image") {
			images = append(images, formFile)
		}
	}

	if arePresetsPresent && len(images) > 0 {
		presetsUploadData := m.mapPresetsFromRequest(r, images)
		filesData = append(filesData, presetsUploadData...)
	}

	return filesData
}

func (m *UploadDataMapper) mapPresetsFromRequest(
	r UploadRequestInterface,
	images []*multipart.FileHeader,
) []*data.UploadFileData {
	presets := r.GetPresets()
	var presetsUploadData []*data.UploadFileData

	for _, image := range images {
		for _, preset := range presets {
			presetFilename := fmt.Sprintf("%d_%s", preset, image.Filename)
			presetUploadData := data.NewUploadFileData(image, presetFilename, r.GetDisk(), r.GetOwnerSign(), preset)
			presetsUploadData = append(presetsUploadData, presetUploadData)
		}
	}

	return presetsUploadData
}
