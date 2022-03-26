package entity

import (
	"database/sql"
	"time"
)

type PrivateFile struct {
	Uuid      string         `json:"uuid" db:"uuid"`
	Path      string         `json:"path" db:"path"`
	Access    string         `json:"access" db:"access"`
	OwnerType sql.NullString `json:"owner_type" db:"owner_type"`
	OwnerUuid sql.NullString `json:"owner_uuid" db:"owner_uuid"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
}
