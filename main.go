// Copyright 2016 The front Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/repejota/logger"
)

// CLI entrypoint.
func main() {
	l := logger.New("default")

	// Start server
	err := Start()
	if err != nil {
		l.Errorf(err.Error())
	}
}
