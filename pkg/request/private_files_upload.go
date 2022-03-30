package request

import "mime/multipart"

type PrivateFilesUploadRequest struct {
	FormFiles []*multipart.FileHeader `form:"files[]" binding:"required"`
	OwnerGuid string                  `form:"owner_guid" binding:"required"`
	OwnerType string                  `form:"owner_type" binding:"required"`
	Presets   []uint                  `form:"presets[]"`
}
