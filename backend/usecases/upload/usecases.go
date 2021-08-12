package usecases

import (
	"log"
	"mime/multipart"
	"path"

	"github.com/google/uuid"
	"github.com/kkhan01/caputo/backend/data/files"
	"github.com/kkhan01/caputo/backend/data/fs"
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

func (usecases usecases) Upload(filename string, file multipart.File) (string, error) {
	saveid := uuid.New()
	savename := saveid.String()

	err := fs.SaveFile(path.Join(usecases.filesRepo.StoragePath, savename), file)
	if err != nil {
		log.Println("Error saving file:", err.Error())
		return savename, err
	}

	err = usecases.filesRepo.WriteFileMeta(saveid, filename)
	if err != nil {
		log.Println("Error writing file meta to db:", err.Error())
		return savename, err
	}

	return savename, nil
}
