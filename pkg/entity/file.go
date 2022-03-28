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

type FileEntity struct {
	Uuid      string    `json:"uuid"`
	Path      string    `json:"path"`
	Access    string    `json:"access"`
	OwnerType string    `json:"owner_type"`
	OwnerUuid string    `json:"owner_uuid"`
	CreatedAt time.Time `json:"created_at"`
}

type File struct {
	Source *multipart.FileHeader `json:"-"`
	Entity *FileEntity
}

func newFileEntity(access string, filename string) *FileEntity {
	uuid := xid.New().String()
	path := storageRoot + "/" + access + "/" + uuid + "_" + filename

	return &FileEntity{
		Uuid:      uuid,
		Access:    access,
		CreatedAt: time.Now(),
		Path:      path,
	}
}

func NewPublicFileEntity(filename string) *FileEntity {
	return newFileEntity(PublicAccessType, filename)
}

func NewPrivateFileEntity(filename string) *FileEntity {
	return newFileEntity(PrivateAccessType, filename)
}

func newFile(access string, source *multipart.FileHeader) *File {
	return &File{
		Source: source,
		Entity: newFileEntity(access, source.Filename),
	}
}

func NewPublicFile(source *multipart.FileHeader) *File {
	return newFile(PublicAccessType, source)
}

func NewPrivateFile(source *multipart.FileHeader, owner FileOwner) *File {
	privateFile := newFile(PrivateAccessType, source)
	privateFile.Entity.OwnerUuid = owner.Uuid
	privateFile.Entity.OwnerType = owner.Type

	return privateFile
}
