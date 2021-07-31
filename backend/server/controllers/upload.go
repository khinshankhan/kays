package controllers

import (
	"fmt"
	"net/http"

	"github.com/kkhan01/caputo/backend/goutils"
	uploadusecases "github.com/kkhan01/caputo/backend/server/usecases/upload"
)

var (
	Size = goutils.MebibyteToBytes(1)
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.MultipartForm)
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

	err = uploadusecases.Upload(filename, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}
