package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"gofile/pkg/entity"
	"gofile/pkg/infrastructure"
)

type FileRepository struct {
	table string
	db    *sqlx.DB
}

func NewFileRepository(db *infrastructure.Database) *FileRepository {
	return &FileRepository{
		table: "files",
		db:    db.DB,
	}
}

func (r *FileRepository) Create(file *entity.File) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (uuid, name, path, access, content_type, type, encoding, owner_guid, owner_type, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		r.table,
	)

	t := r.db.MustBegin()
	t.MustExec(query, file.Uuid, file.Name, file.Path, file.Access, file.ContentType, file.Type, file.Encoding, file.OwnerGuid, file.OwnerType, file.CreatedAt.String())

	return t.Commit()
}

func (r *FileRepository) FindByUuidAndAccess(uuid string, access string) (*entity.File, error) {
	file := entity.File{}

	query := fmt.Sprintf("SELECT * FROM %s WHERE uuid = $1 AND access = $2", r.table)
	err := r.db.Get(&file, query, uuid, access)
	if file.Uuid != uuid {
		return nil, errors.New("file not found")
	}

	return &file, err
}

func (r *FileRepository) DeleteByUuid(uuid string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE uuid = $1", r.table)
	_, err := r.db.Exec(query, uuid)

	return err
}
