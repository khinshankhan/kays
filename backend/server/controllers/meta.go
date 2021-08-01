package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kkhan01/caputo/backend/server/config"
	metausecases "github.com/kkhan01/caputo/backend/server/usecases/meta"
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
