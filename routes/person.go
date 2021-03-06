// Copyright 2016 The front Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/repejota/logger"
	"github.com/tvtio/front/catalog"
	"github.com/tvtio/tmdb"
)

// Person is the /movie/:id route
func Person(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	l := logger.New("default")

	id := ps.ByName("id")
	person, err := catalog.Person(id)
	if err != nil {
		l.Errorf(err.Error())
		http.Error(w, "HTTP 500 : Internal Server Error", 500)
		return
	}
	context := struct {
		Title  string
		Person tmdb.Person
	}{
		"tvt.io",
		person,
	}
	t := template.Must(template.ParseFiles(
		"templates/person.html",
		"templates/partials/facebook.html",
		"templates/partials/footer.html",
		"templates/partials/javascript.html",
		"templates/partials/css.html",
	))
	t.Execute(w, context)
}
