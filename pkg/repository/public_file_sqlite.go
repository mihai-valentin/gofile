package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"gofile/pkg/entity"
)

type PublicFileSqlite struct {
	db    *sqlx.DB
	table string
}

func NewPublicFileSqlite(db *sqlx.DB) *PublicFileSqlite {
	return &PublicFileSqlite{
		db:    db,
		table: "files",
	}
}

func (r *PublicFileSqlite) CreatePublicFile(file *entity.File) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (uuid, name, path, access, content_type, type, encoding, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		r.table,
	)

	t := r.db.MustBegin()
	t.MustExec(query, file.Uuid, file.Name, file.Path, file.Access, file.ContentType, file.Type, file.Encoding, file.CreatedAt.String())

	return t.Commit()
}

func (r *PublicFileSqlite) FindPublicFileByUuid(uuid string) (*entity.File, error) {
	file := entity.File{}

	query := fmt.Sprintf("SELECT * FROM %s WHERE uuid = $1 AND access = $2", r.table)
	// TODO: extract hardcode
	err := r.db.Get(&file, query, uuid, "public")
	if file.Uuid != uuid {
		return nil, errors.New("file not found")
	}

	return &file, err
}

func (r *PublicFileSqlite) DeletePublicFileByUuid(uuid string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE uuid = $1 AND access = $2", r.table)
	// TODO: extract hardcode
	_, err := r.db.Exec(query, uuid, "public")

	return err
}
