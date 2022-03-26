package entity

import (
	"mime/multipart"
)

type PrivateFilesList struct {
	FileOwner FileOwner
	Payload   []*multipart.FileHeader `form:"payload[]" binding:"required"`
}
