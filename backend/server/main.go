package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// Version and BuildData get replaced during build with the commit hash and time
// of build.
var (
	CommitHash = ""
	BuildDate  = ""
)

// Meta holds project metadata
type Meta struct {
	CommitHash string    `json:"commitHash"`
	BuildDate  string    `json:"buildDate"`
	SystemTime time.Time `json:"systemTime"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", meta)

	port := 5000
	if p := os.Getenv("PORT"); p != "" {
		_, err := fmt.Sscanf(p, "%d", &port)
		if err != nil {
			log.Fatal("Invalid PORT in env")
		}
	}

	fmt.Printf("Listening on port %d\n", port)
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%d", port), router),
	)
}

func meta(w http.ResponseWriter, r *http.Request) {
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
