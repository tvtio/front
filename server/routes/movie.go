// Copyright 2016 The front Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/mgutz/ansi"
	"github.com/repejota/logger"
	"github.com/tvtio/front/catalog"
	"github.com/tvtio/tmdb"
)

// Movie is the /movie/:id route
func Movie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	l := logger.New("default")

	id := ps.ByName("id")
	movie, err := catalog.Movie(id)
	if err != nil {
		l.Errorf(ansi.Color("FATAL: ", "red"), err)
	}
	context := struct {
		Title string
		Movie tmdb.Movie
	}{
		"tvt.io",
		movie,
	}
	t := template.Must(template.ParseFiles(
		"templates/movie.html",
		"templates/partials/facebook.html",
		"templates/partials/footer.html",
		"templates/partials/javascript.html",
		"templates/partials/css.html",
	))
	t.Execute(w, context)
}
