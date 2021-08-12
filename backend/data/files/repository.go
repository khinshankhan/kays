package files

import (
	"database/sql"

	"github.com/google/uuid"
)

type (
	Repository struct {
		db          *sql.DB
		StoragePath string
	}
)

func NewRepository(db *sql.DB, storagePath string) Repository {
	return Repository{
		db:          db,
		StoragePath: storagePath,
	}
}

func (repo Repository) WriteFileMeta(saveid uuid.UUID, filename string) error {
	_, err := repo.db.Exec(
		"INSERT INTO files (saveid, filename) VALUES ($1, $2);",
		saveid,
		filename,
	)

	return err
}
