package entity

import (
	"github.com/rs/xid"
	"mime/multipart"
	"time"
)

const (
	storageRoot       = "storage"
	PublicAccessType  = "public"
	PrivateAccessType = "private"
)

type File struct {
	Source    *multipart.FileHeader `json:"-"`
	Uuid      string                `json:"uuid"`
	Path      string                `json:"path"`
	Access    string                `json:"access"`
	OwnerType string                `json:"owner_type"`
	OwnerUuid string                `json:"owner_uuid"`
	CreatedAt time.Time             `json:"created_at"`
}

func newFile(access string, source *multipart.FileHeader) *File {
	uuid := xid.New().String()
	path := storageRoot + "/" + access + "/" + uuid + "_" + source.Filename

	return &File{
		Source:    source,
		Uuid:      uuid,
		Access:    access,
		CreatedAt: time.Now(),
		Path:      path,
	}
}

func NewPublicFile(source *multipart.FileHeader) *File {
	return newFile(PublicAccessType, source)
}

func NewPrivateFile(source *multipart.FileHeader, owner FileOwner) *File {
	privateFile := newFile(PrivateAccessType, source)
	privateFile.OwnerUuid = owner.Uuid
	privateFile.OwnerType = owner.Type

	return privateFile
}
