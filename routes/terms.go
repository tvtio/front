// Copyright 2016 The front Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/repejota/logger"
)

// Terms is the /terms route
func Terms(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	l := logger.New("default")

	t, err := template.ParseFiles(
		"templates/terms.html",
		"templates/partials/facebook.html",
		"templates/partials/footer.html",
		"templates/partials/javascript.html",
		"templates/partials/css.html",
	)
	if err != nil {
		l.Errorf(err.Error())
		http.Error(w, "HTTP 500 : Internal Server Error", 500)
		return
	}
	context := struct {
		Title string
	}{
		"tvt.io",
	}
	t.Execute(w, context)
}
