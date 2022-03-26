package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gofile/pkg/entity"
)

type PrivateFileRepository struct {
	db *sqlx.DB
}

func (r *PrivateFileRepository) Create(file entity.File) {
	query := fmt.Sprintf(
		"INSERT INTO %s (uuid, path, access, owner_uuid, owner_type, created_at) VALUES ($1, $2, $3, $4, $5, $6)",
		filesTable,
	)
	fmt.Println(query)

	t := r.db.MustBegin()
	t.MustExec(query, file.Uuid, file.Path, file.Access, file.OwnerUuid, file.OwnerType, file.CreatedAt.String())
	t.Commit()
}

func (r *PrivateFileRepository) FindByUuid(uuid string) (entity.PrivateFile, error) {
	file := entity.PrivateFile{}

	query := fmt.Sprintf("SELECT * FROM %s WHERE access = $1 AND uuid = $2", filesTable)
	err := r.db.Get(&file, query, entity.PrivateAccessType, uuid)

	return file, err
}

func (r *PrivateFileRepository) DeleteByUuid(uuid string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE access = $2 AND uuid = $1", filesTable)
	_, err := r.db.Exec(query, entity.PrivateAccessType, uuid)

	return err
}

func NewPrivateFileRepository(db *sqlx.DB) *PrivateFileRepository {
	return &PrivateFileRepository{db: db}
}
