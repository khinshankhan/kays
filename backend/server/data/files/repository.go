package files

import (
	"database/sql"
)

type (
	Repository struct {
		db *sql.DB
	}
)

func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

func (repo Repository) WriteFileMeta(savename, filename string) error {
	_, err := repo.db.Exec(
		"INSERT INTO files (name, original) VALUES ($1, $2);",
		savename,
		filename,
	)

	return err
}
