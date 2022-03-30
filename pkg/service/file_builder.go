package service

import (
	"fmt"
	"github.com/rs/xid"
	"gofile/pkg/entity"
	"strings"
	"time"
)

type FileBuilder struct {
}

func (s *FileBuilder) buildFile(access string, filename string, contentType string) *entity.File {
	uuid := xid.New().String()
	path := "storage/" + access + "/" + uuid + "_" + filename
	contentTypeParts := strings.Split(contentType, "/")

	return &entity.File{
		Uuid:        uuid,
		Name:        filename,
		Access:      access,
		ContentType: contentType,
		Type:        contentTypeParts[0],
		Encoding:    contentTypeParts[1],
		Path:        path,
		CreatedAt:   time.Now(),
	}
}

func (s *FileBuilder) buildPresetFile(original *entity.File, scale uint) *entity.File {
	uuid := xid.New().String()
	filename := fmt.Sprintf("%d_%s", scale, original.Name)
	path := fmt.Sprintf("storage/%s/%s_%s", original.Access, original.Uuid, filename)

	return &entity.File{
		Uuid:        uuid,
		Name:        filename,
		Access:      original.Access,
		ContentType: original.ContentType,
		Type:        original.Type,
		Encoding:    original.Encoding,
		Path:        path,
		CreatedAt:   time.Now(),
	}
}
