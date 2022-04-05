package data

import (
	"github.com/rs/xid"
	"mime/multipart"
	"strings"
)

type FileUpload struct {
	Source    *multipart.FileHeader
	Uuid      string
	Filename  string
	Disk      string
	Path      string
	OwnerSign string
	Scale     uint
}

func NewFileUpload(
	source *multipart.FileHeader,
	filename string,
	disk string,
	ownerSign string,
	scale uint,
) *FileUpload {
	uuid := xid.New().String()

	return &FileUpload{
		Source:    source,
		Uuid:      uuid,
		Filename:  filename,
		Disk:      disk,
		Path:      disk + "/" + uuid + "_" + filename,
		OwnerSign: ownerSign,
		Scale:     scale,
	}
}

func (d *FileUpload) GetSource() *multipart.FileHeader {
	return d.Source
}

func (d *FileUpload) GetUuid() string {
	return d.Uuid
}

func (d *FileUpload) GetFilename() string {
	return d.Filename
}

func (d *FileUpload) GetDisk() string {
	return d.Disk
}

func (d *FileUpload) GetPath() string {
	return d.Path
}

func (d *FileUpload) GetOwnerSign() string {
	return d.OwnerSign
}

func (d *FileUpload) GetScale() uint {
	return d.Scale
}

func (d *FileUpload) IsResizable() bool {
	return d.Scale != 0
}

func (d *FileUpload) GetEncoding() string {
	return strings.Split(d.Source.Header.Get("Content-Type"), "/")[1]
}

func (d *FileUpload) IsImage() bool {
	return strings.HasPrefix(d.Source.Header.Get("Content-Type"), "image")
}
