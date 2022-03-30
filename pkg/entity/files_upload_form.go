package entity

import "mime/multipart"

type FilesUploadForm struct {
	FormFiles []*multipart.FileHeader `form:"files[]" binding:"required"`
	Presets   []uint                  `form:"presets[]"`
	OwnerGuid string                  `form:"owner_guid"`
	OwnerType string                  `form:"owner_type"`
}
