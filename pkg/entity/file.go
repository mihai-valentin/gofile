package entity

import (
	"database/sql"
	"time"
)

type File struct {
	Uuid        string         `json:"uuid" db:"uuid"`
	Name        string         `json:"filename" db:"name"`
	Access      string         `json:"access" db:"access"`
	ContentType string         `json:"-" db:"content_type"`
	Type        string         `json:"-" db:"type"`
	Encoding    string         `json:"-" db:"encoding"`
	Path        string         `json:"path" db:"path"`
	OwnerType   sql.NullString `json:"owner_type" db:"owner_type"`
	OwnerGuid   sql.NullString `json:"owner_guid" db:"owner_guid"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
}

func (f *File) IsOfType(contentType string) bool {
	return f.Type == contentType
}

func (f *File) IsAnImage() bool {
	return f.IsOfType("image")
}
