package entity

import "time"

type File struct {
	Uuid      string    `json:"-" db:"uuid"`
	Name      string    `json:"-" db:"name"`
	Disk      string    `json:"-" db:"disk"`
	Path      string    `json:"-" db:"path"`
	OwnerSign string    `json:"-" db:"owner_sign"`
	CreatedAt time.Time `json:"-" db:"created_at"`
}

func (e *File) GetUuid() string {
	return e.Uuid
}

func (e *File) GetName() string {
	return e.Name
}

func (e *File) GetDisk() string {
	return e.Disk
}

func (e *File) GetPath() string {
	return e.Path
}

func (e *File) GetOwnerSign() string {
	return e.OwnerSign
}

func (e *File) GetCreatedAt() string {
	return e.CreatedAt.String()
}
