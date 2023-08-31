package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Set up Loki configuration
	lokiHook, err := loki.NewLokiHook(
		"http://loki-server-url:3100/loki/api/v1/push",
		"my-service-name",
		log.InfoLevel,
		&http.Client{}, // You might need to import "net/http" and pass a valid client.
	)
	if err != nil {
		fmt.Println("Error creating Loki hook:", err)
		os.Exit(1)
	}

	// Add Loki hook to logrus
	log.AddHook(lokiHook)

	// Start logging
	log.Info("This is an info log message.")
	log.Error("This is an error log message.")
}
