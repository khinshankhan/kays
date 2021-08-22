package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/kkhan01/caputo/backend/data/files"
)

func ServeHandler(filesRepo files.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			return
		}

		// get UUID from request
		vars := mux.Vars(r)
		fileUUID, ok := vars["fileUUID"]
		if !ok {
			http.Error(w, "No file uuid provided", http.StatusBadRequest)
			return
		}

		// static serve the file if it exists
		filePath := path.Join(filesRepo.StoragePath, fileUUID)
		http.ServeFile(w, r, filePath)
	}
}

func FilesMetaHandler(filesRepo files.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			return
		}

		files, err := filesRepo.GetFilesMeta()
		if err != nil {
			msg := "Could not retrieve files list:"
			log.Println(msg, err)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		filesJSON, err := json.Marshal(files)
		if err != nil {
			msg := "Failed to encode files list:"
			log.Println(msg, err)
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(filesJSON)
	}
}
