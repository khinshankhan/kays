package files

import (
	"database/sql"
	"fmt"
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

func (repo Repository) WriteFileMeta(saveename, filename string) error {
	fmt.Println("Write file meta.")
	return nil
}
