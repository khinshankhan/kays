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

	FileMeta struct {
		SaveID   uuid.UUID `json:"uuid"`
		Filename string    `json:"filename"`
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

func (repo Repository) GetFilesMeta() ([]FileMeta, error) {
	rows, err := repo.db.Query("SELECT * FROM files;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]FileMeta, 0)
	for rows.Next() {
		var result FileMeta
		err := rows.Scan(&result.SaveID, &result.Filename)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}
