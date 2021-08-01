package usecases

import (
	"log"
	"mime/multipart"

	"github.com/kkhan01/caputo/backend/server/data/files"
	"github.com/kkhan01/caputo/backend/server/data/fs"
)

func Upload(repo files.Repository, filename string, file multipart.File) error {
	err := fs.SaveFile(filename, file)
	if err != nil {
		log.Println("Error saving file:", err.Error())
		return err
	}

	err = repo.WriteFileMeta(filename, filename)
	if err != nil {
		log.Println("Error writing file meta to db:", err.Error())
		return err
	}

	return nil
}
