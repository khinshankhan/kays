package usecases

import (
	"log"
	"mime/multipart"

	"github.com/kkhan01/caputo/backend/server/data/files"
	"github.com/kkhan01/caputo/backend/server/data/fs"
)

type (
	// Usecases declares available services
	Usecases interface {
		Upload(filename string, file multipart.File) error
	}

	// usecases declares the dependencies for the service
	usecases struct {
		filesRepo files.Repository
	}
)

// NewUsecases returns Usecases
func NewUsecases(filesRepo files.Repository) usecases {
	return usecases{
		filesRepo: filesRepo,
	}
}

func (usecases usecases) Upload(filename string, file multipart.File) error {
	err := fs.SaveFile(filename, file)
	if err != nil {
		log.Println("Error saving file:", err.Error())
		return err
	}

	err = usecases.filesRepo.WriteFileMeta(filename, filename)
	if err != nil {
		log.Println("Error writing file meta to db:", err.Error())
		return err
	}

	return nil
}
