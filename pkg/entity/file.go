package entity

import "time"

type File struct {
	Uuid      string    `json:"uuid" db:"uuid"`
	Name      string    `json:"filename" db:"name"`
	Disk      string    `json:"disk" db:"disk"`
	Path      string    `json:"-" db:"path"`
	OwnerSign string    `json:"-" db:"owner_sign"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
