package entity

import (
	"time"
)

type PublicFile struct {
	Uuid      string    `json:"uuid" db:"uuid"`
	Path      string    `json:"path" db:"path"`
	Access    string    `json:"access" db:"access"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
