package request

import "mime/multipart"

type PubicFilesUploadRequest struct {
	FormFiles []*multipart.FileHeader `form:"files[]" binding:"required"`
	Presets   []uint                  `form:"presets[]"`
}
