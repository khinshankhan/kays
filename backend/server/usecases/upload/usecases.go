package usecases

import (
	"mime/multipart"

	"github.com/kkhan01/caputo/backend/server/data/files"
	"github.com/kkhan01/caputo/backend/server/data/fs"
)

func Upload(repo files.Repository, filename string, file multipart.File) error {
	err := fs.SaveFile(filename, file)
	if err != nil {
		return err
	}

	return repo.WriteFileMeta(filename, filename)
}
