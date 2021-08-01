package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kkhan01/caputo/backend/config"
	metausecases "github.com/kkhan01/caputo/backend/usecases/meta"
)

func MetaHandler(metaConfig *config.MetaConfig) http.HandlerFunc {
	metaService := metausecases.NewUsecases(metaConfig)

	return func(w http.ResponseWriter, r *http.Request) {
		metaInfo := metaService.GetMeta()

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(metaInfo); err != nil {
			log.Println(err.Error())
		}
	}
}
