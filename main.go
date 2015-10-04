package main

import (
	"flag"
	"log"

	"github.com/tvtio/front/config"
	"github.com/tvtio/front/server"
)

func main() {
	// Flags
	filename := flag.String("config", "./config.json", "Configuration file")
	flag.Parse()

	// Load configuration
	configuration, err := config.Load(*filename)
	if err != nil {
		log.Fatal(err)
	}

	// Start server
	log.Fatal(server.Start(configuration))
}
