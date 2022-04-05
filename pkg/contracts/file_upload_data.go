package contracts

import "mime/multipart"

type FileUploadDataInterface interface {
	GetSource() *multipart.FileHeader
	GetUuid() string
	GetFilename() string
	GetDisk() string
	GetOwnerSign() string
	GetPath() string
	GetScale() uint
	IsImage() bool
	IsResizable() bool
	GetEncoding() string
}
