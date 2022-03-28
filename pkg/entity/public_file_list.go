package entity

import (
	"mime/multipart"
)

type PublicFilesList struct {
	FormFiles []*multipart.FileHeader `form:"files[]" binding:"required"`
	Presets   []uint                  `form:"presets[]"`
}
