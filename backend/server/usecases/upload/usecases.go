package usecases

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func Upload(filename string, file multipart.File) error {
	// Create file
	dst, err := os.Create(filename)
	defer dst.Close()
	if err != nil {
		return err
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		return err
	}

	fmt.Println("Saved", filename)

	return nil
}
