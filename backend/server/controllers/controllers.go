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
func Handle(cfg *config.Config) {
	router := CreateRouter(cfg)

	address := cfg.Web.SocketAddress()
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
func CreateRouter(cfg *config.Config) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/upload", UploadHandler)
	router.HandleFunc("/meta", MetaHandler(cfg.Meta))

	return router
}
