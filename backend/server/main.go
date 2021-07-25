package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/kkhan01/caputo/backend/server/routes"
)

// Version and BuildData get replaced during build with the commit hash and time
// of build.
var (
	CommitHash = ""
	BuildDate  = ""
)

func main() {
	port := loadConfiguration()

	start(port)
	// Wait for termination signal
	signalChannel := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChannel
		log.Println("Received termination signal, attempting to gracefully shut down")
		stop()
		// TODO: add in graceful services shutdown
		done <- true
	}()
	<-done
	log.Println("Shutting down")
}

func start(port int) {
	go routes.Handle(CommitHash, BuildDate, port)
}

func stop() {
	routes.Shutdown()
}

// TODO: build this out to read in a full configuration
func loadConfiguration() int {
	var err error

	port := 5000
	if p := os.Getenv("PORT"); p != "" {
		if port, err = strconv.Atoi(p); err != nil {
			log.Fatal("Invalid PORT in env")
		}
	}

	return port
}
