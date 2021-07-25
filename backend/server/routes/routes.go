package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	CommitHash = ""
	BuildDate  = ""
	server     *http.Server
)

// Handle creates the router and starts the server
func Handle(commitHash, buildDate string, port int) {
	CommitHash = commitHash
	BuildDate = buildDate

	router := CreateRouter()

	server = &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	fmt.Printf("Listening on port %d\n", port)
	server.ListenAndServe()
}

// Shutdown stops the server
func Shutdown() {
	if server != nil {
		_ = server.Shutdown(context.TODO())
		server = nil
	}
}

// CreateRouter creates the router for the http server
func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/meta", MetaHandler)

	return router
}
