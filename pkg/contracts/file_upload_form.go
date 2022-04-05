package contracts

import "mime/multipart"

type FileUploadFormInterface interface {
	GetFormFiles() []*multipart.FileHeader
	GetPresets() []uint
	GetDisk() string
	GetOwnerSign() string
}
