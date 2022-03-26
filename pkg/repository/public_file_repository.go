package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gofile/pkg/entity"
)

type PublicFileRepository struct {
	db *sqlx.DB
}

func (r *PublicFileRepository) Create(file entity.File) {
	query := fmt.Sprintf(
		"INSERT INTO %s (uuid, path, access, created_at) VALUES ($1, $2, $3, $4)",
		filesTable,
	)

	t := r.db.MustBegin()
	t.MustExec(query, file.Uuid, file.Path, file.Access, file.CreatedAt.String())
	t.Commit()
}

func (r *PublicFileRepository) FindByUuid(uuid string) (entity.PublicFile, error) {
	file := entity.PublicFile{}

	query := fmt.Sprintf("SELECT uuid, path, access, created_at FROM %s WHERE access = $1 AND uuid = $2", filesTable)
	err := r.db.Get(&file, query, entity.PublicAccessType, uuid)

	return file, err
}

func (r *PublicFileRepository) DeleteByUuid(uuid string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE access = $2 AND uuid = $1", filesTable)
	_, err := r.db.Exec(query, entity.PublicAccessType, uuid)

	return err
}

func NewPublicFileRepository(db *sqlx.DB) *PublicFileRepository {
	return &PublicFileRepository{db: db}
}
