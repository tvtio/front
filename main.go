package main

import (
	"github.com/tvtio/front/logger"
	"github.com/tvtio/front/server"
)

func main() {
	// Start server
	err := server.Start()
	if err != nil {
		logger.Fatal(err.Error())
	}
}
