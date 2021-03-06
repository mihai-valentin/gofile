package mapper

import (
	"fmt"
	"gofile/pkg/contracts"
	"gofile/pkg/data"
	"mime/multipart"
	"strings"
)

type UploadData struct {
}

func (m *UploadData) MapFromRequest(r contracts.FileUploadFormInterface) []contracts.FileUploadDataInterface {
	formFiles := r.GetFormFiles()
	arePresetsPresent := len(r.GetPresets()) > 0

	var images []*multipart.FileHeader
	var filesData []contracts.FileUploadDataInterface

	for _, formFile := range formFiles {
		uploadFileData := data.NewFileUpload(formFile, formFile.Filename, r.GetDisk(), r.GetOwnerSign(), 0)
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

func (m *UploadData) mapPresetsFromRequest(
	r contracts.FileUploadFormInterface,
	images []*multipart.FileHeader,
) []contracts.FileUploadDataInterface {
	presets := r.GetPresets()
	var presetsUploadData []contracts.FileUploadDataInterface

	for _, image := range images {
		for _, preset := range presets {
			presetFilename := fmt.Sprintf("%d_%s", preset, image.Filename)
			presetUploadData := data.NewFileUpload(image, presetFilename, r.GetDisk(), r.GetOwnerSign(), preset)
			presetsUploadData = append(presetsUploadData, presetUploadData)
		}
	}

	return presetsUploadData
}
