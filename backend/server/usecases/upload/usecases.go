package usecases

import (
	"mime/multipart"

	"github.com/kkhan01/caputo/backend/server/data/fs"
)

func Upload(filename string, file multipart.File) error {
	return fs.SaveFile(filename, file)
}
