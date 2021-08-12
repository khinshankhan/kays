package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kkhan01/caputo/backend/data/files"
)

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
