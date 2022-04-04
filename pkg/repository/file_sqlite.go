package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gofile/pkg/entity"
	"time"
)

type FileDataInterface interface {
	GetUuid() string
	GetFilename() string
	GetDisk() string
	GetPath() string
	GetOwnerSign() string
}

type FileSqlite struct {
	db    *sqlx.DB
	table string
}

func NewPublicFileSqlite(db *sqlx.DB) *FileSqlite {
	return &FileSqlite{
		db:    db,
		table: "files",
	}
}

func (r FileSqlite) buildEntity(fileData FileDataInterface) *entity.File {
	return &entity.File{
		Uuid:      fileData.GetUuid(),
		Name:      fileData.GetFilename(),
		Disk:      fileData.GetDisk(),
		Path:      fileData.GetPath(),
		OwnerSign: fileData.GetOwnerSign(),
		CreatedAt: time.Now(),
	}
}

func (r *FileSqlite) CreateFile(fileData FileDataInterface) (*entity.File, error) {
	file := r.buildEntity(fileData)

	query := fmt.Sprintf(
		"INSERT INTO %s (uuid, name, disk, path, owner_sign, created_at) VALUES ($1, $2, $3, $4, $5, $6)",
		r.table,
	)

	t := r.db.MustBegin()
	t.MustExec(query, file.Uuid, file.Name, file.Disk, file.Path, file.OwnerSign, file.CreatedAt)

	if err := t.Commit(); err != nil {
		return nil, err
	}

	return file, nil
}

func (r *FileSqlite) FindByUuid(uuid string) (*entity.File, error) {
	file := entity.File{}

	query := fmt.Sprintf("SELECT * FROM %s WHERE uuid = $1", r.table)
	err := r.db.Get(&file, query, uuid)

	return &file, err
}

func (r *FileSqlite) DeleteByUuid(uuid string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE uuid = $1", r.table)
	_, err := r.db.Exec(query, uuid)

	return err
}
