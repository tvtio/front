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

// Index is the / route
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	l := logger.New("default")

	// Get popular movies
	popularMovies, err := catalog.PopularMovies()
	if err != nil {
		l.Errorf(err.Error())
	}

	// Get popular tv
	popularTV, err := catalog.PopularTV()
	if err != nil {
		l.Errorf(err.Error())
	}

	// Build template
	t := template.Must(template.ParseFiles(
		"templates/index.html",
		"templates/partials/facebook.html",
		"templates/partials/footer.html",
		"templates/partials/javascript.html",
		"templates/partials/css.html",
	))

	// Template context
	context := struct {
		Title         string
		PopularMovies tmdb.SearchMovieResult
		PopularTV     tmdb.SearchTVResult
		BG1           string
	}{
		"tvt.io",
		popularMovies,
		popularTV,
		popularMovies.Results[0].BackdropPath,
	}

	// Render template
	t.Execute(w, context)
}
