package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kkhan01/caputo/backend/server/config"
)

var (
	server *http.Server
)

// Handle creates the router and starts the server
func Handle(metaConfig *config.MetaConfig, webConfig *config.WebConfig) {
	router := CreateRouter(metaConfig)

	address := webConfig.SocketAddress()
	server = &http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	fmt.Printf("Listening on %s\n", address)
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
func CreateRouter(metaConfig *config.MetaConfig) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/meta", MetaHandler(metaConfig))
	router.HandleFunc("/upload", UploadHandler)

	return router
}
