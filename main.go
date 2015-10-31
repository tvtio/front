package main

import (
	"flag"

	"github.com/tvtio/front/config"
	"github.com/tvtio/front/logger"
	"github.com/tvtio/front/server"
)

func main() {
	// Flags
	filename := flag.String("config", "./config.json", "Configuration file")
	flag.Parse()

	// Load configuration
	configuration, err := config.Load(*filename)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Start server
	err = server.Start(configuration)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
