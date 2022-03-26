package entity

import (
	"mime/multipart"
)

type PublicFilesList struct {
	Payload []*multipart.FileHeader `form:"payload[]" binding:"required"`
}
