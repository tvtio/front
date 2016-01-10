package main

import (
	"github.com/repejota/logger"
	"github.com/tvtio/front/server"
)

// CLI entrypoint.
func main() {
	l := logger.New("default")

	// Start server
	err := server.Start()
	if err != nil {
		l.Errorf(err.Error())
	}
}
