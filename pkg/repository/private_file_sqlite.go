package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gofile/pkg/entity"
)

type PrivateFileSqlite struct {
	db    *sqlx.DB
	table string
}

func NewPrivateFileSqlite(db *sqlx.DB) *PrivateFileSqlite {
	return &PrivateFileSqlite{
		db:    db,
		table: "files",
	}
}

func (r *PrivateFileSqlite) CreatePrivateFile(file *entity.File) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (uuid, name, access, content_type, type, encoding, path, owner_guid, owner_type, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		r.table,
	)
	fmt.Println(query)

	t := r.db.MustBegin()
	t.MustExec(query, file.Uuid, file.Name, file.Access, file.ContentType, file.Type, file.Encoding, file.Path, file.OwnerGuid, file.OwnerType, file.CreatedAt.String())

	return t.Commit()
}

func (r *PrivateFileSqlite) FindPrivateFileByUuidAndOwner(uuid string, ownerGuid string, ownerType string) (*entity.File, error) {
	file := new(entity.File)

	query := fmt.Sprintf("SELECT * FROM %s WHERE access = $1 AND uuid = $2 AND owner_guid = $3 AND owner_type = $4", r.table)
	err := r.db.Get(&file, query, "private", uuid, ownerGuid, ownerType)

	return file, err
}

func (r *PrivateFileSqlite) DeletePrivateFileByUuidAndOwner(uuid string, ownerGuid string, ownerType string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE  uuid = $1 AND access = $2 AND owner_guid = $3 AND owner_type = $4", r.table)
	_, err := r.db.Exec(query, "private", uuid, ownerGuid, ownerType)

	return err
}
