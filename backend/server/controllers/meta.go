package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/kkhan01/caputo/backend/server/config"
)

// Meta holds project metadata
type Meta struct {
	CommitHash string    `json:"commitHash"`
	BuildDate  string    `json:"buildDate"`
	SystemTime time.Time `json:"systemTime"`
}

func MetaHandler(metaConfig *config.MetaConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentTimestamp := time.Now().UTC()

		metainfo := Meta{
			BuildDate:  metaConfig.BuildDate,
			CommitHash: metaConfig.CommitHash,
			SystemTime: currentTimestamp,
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(metainfo); err != nil {
			log.Println(err.Error())
		}
	}
}
