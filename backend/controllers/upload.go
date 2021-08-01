package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kkhan01/caputo/backend/data/files"
	uploaddomain "github.com/kkhan01/caputo/backend/domain/upload"
	uploadusecases "github.com/kkhan01/caputo/backend/usecases/upload"
	"github.com/kkhan01/caputo/backend/utils"
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

		savename, err := uploadService.Upload(filename, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		uploadInfo := uploaddomain.Uploaded{
			Filename: savename,
			Status:   http.StatusCreated,
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(uploadInfo); err != nil {
			log.Println(err.Error())
		}
	}
}
