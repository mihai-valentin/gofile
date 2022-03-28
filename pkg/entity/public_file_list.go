package entity

import (
	"mime/multipart"
)

type PublicFilesList struct {
	Files  []*multipart.FileHeader `form:"files[]" binding:"required"`
	Width  []uint16                `form:"presets[width][]"`
	Height []uint16                `form:"presets[height][]"`
}
