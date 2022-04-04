package data

import (
	"github.com/rs/xid"
	"mime/multipart"
	"strings"
)

type UploadFileData struct {
	Source    *multipart.FileHeader
	Uuid      string
	Filename  string
	Disk      string
	Path      string
	OwnerSign string
	Scale     uint
}

func NewUploadFileData(
	source *multipart.FileHeader,
	filename string,
	disk string,
	ownerSign string,
	scale uint,
) *UploadFileData {
	uuid := xid.New().String()

	return &UploadFileData{
		Source:    source,
		Uuid:      uuid,
		Filename:  filename,
		Disk:      disk,
		Path:      disk + "/" + uuid + "_" + filename,
		OwnerSign: ownerSign,
		Scale:     scale,
	}
}

func (d *UploadFileData) GetSource() *multipart.FileHeader {
	return d.Source
}

func (d *UploadFileData) GetUuid() string {
	return d.Uuid
}

func (d *UploadFileData) GetFilename() string {
	return d.Filename
}

func (d *UploadFileData) GetDisk() string {
	return d.Disk
}

func (d *UploadFileData) GetPath() string {
	return d.Path
}

func (d *UploadFileData) GetOwnerSign() string {
	return d.OwnerSign
}

func (d *UploadFileData) GetScale() uint {
	return d.Scale
}

func (d *UploadFileData) GetEncoding() string {
	return strings.Split(d.Source.Header.Get("Content-Type"), "/")[1]
}

func (d *UploadFileData) IsResizable() bool {
	return d.Scale != 0
}

func (d *UploadFileData) IsImage() bool {
	return strings.HasPrefix(d.Source.Header.Get("Content-Type"), "image")
}
