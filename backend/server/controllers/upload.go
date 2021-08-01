package controllers

import (
	"fmt"
	"net/http"

	"github.com/kkhan01/caputo/backend/server/data/files"
	uploadusecases "github.com/kkhan01/caputo/backend/server/usecases/upload"
	"github.com/kkhan01/caputo/backend/server/utils"
)

var (
	// Allow uploads up to 10 mb
	Size = utils.MebibyteToBytes(10)
)

func UploadHandler(filesRepo files.Repository) http.HandlerFunc {
	uploadService := uploadusecases.NewUsecases(filesRepo)

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			return
		}

		r.ParseMultipartForm(Size)

		// Get handler for filename, size and headers
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		defer file.Close()

		if handler.Size > Size {
			fmt.Println("Size too big:", handler.Size)
			return
		}

		filename := r.FormValue("name")
		if filename == "" {
			filename = handler.Filename
		}

		err = uploadService.Upload(filename, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Fprintf(w, "Successfully Uploaded File\n")
	}
}
