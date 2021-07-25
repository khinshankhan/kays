package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Meta holds project metadata
type Meta struct {
	CommitHash string    `json:"commitHash"`
	BuildDate  string    `json:"buildDate"`
	SystemTime time.Time `json:"systemTime"`
}

func MetaHandler(w http.ResponseWriter, r *http.Request) {

	currentTimestamp := time.Now().UTC()

	metainfo := Meta{
		BuildDate:  BuildDate,
		CommitHash: CommitHash,
		SystemTime: currentTimestamp,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(metainfo); err != nil {
		log.Println(err.Error())
	}
}
