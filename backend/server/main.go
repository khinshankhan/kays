package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kkhan01/caputo/backend/server/config"
	"github.com/kkhan01/caputo/backend/server/controllers"
)

// Version and BuildData get replaced during build with the commit hash and time
// of build.
var (
	CommitHash = ""
	BuildDate  = ""
)

var (
	envFile = flag.String("env", "", "Path to env file in either yaml or json format")
)

func main() {
	flag.Parse()

	cfg, err := loadConfiguration()
	if err != nil {
		panic(err)
	}

	start(cfg)
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

func start(cfg *config.Config) {
	go controllers.Handle(cfg.Meta, cfg.Web)
}

func stop() {
	controllers.Shutdown()
}

// TODO: build this out to read in a full configuration
func loadConfiguration() (*config.Config, error) {
	return config.Load(CommitHash, BuildDate, *envFile)
}
