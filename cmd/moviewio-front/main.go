// Copyright 2014-2015 Moview.io authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"html/template"
	"net/http"
)

// defaultHandler
//
// Handles the default route '/'
// TODO:
//	* templatePath should be loaded from a configuration file
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	templateBasePath := "./src/github.com/moviewio/front/assets/templates"
	templatePath := templateBasePath + "/index.html"
	t, _ := template.ParseFiles(templatePath)
	t.Execute(w, nil)
}

// main
//
// Default entry point
// TODO:
//	* Load port and address from a configuration file
func main() {
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":9000", nil)
}
