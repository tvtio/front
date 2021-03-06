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

// Episode is the /tv/:id/season/:sid/episode/:enumber route
func Episode(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	l := logger.New("default")

	id := params.ByName("id")
	tv, err := catalog.TV(id)
	if err != nil {
		l.Errorf(err.Error())
		http.Error(w, "HTTP 500 : Internal Server Error", 500)
		return
	}
	snumber := params.ByName("snumber")
	season, err := catalog.Season(id, snumber)
	if err != nil {
		logger.New("default").Errorf(err.Error())
		http.Error(w, "HTTP 500 : Internal Server Error", 500)
		return
	}
	enumber := params.ByName("enumber")
	episode, err := catalog.Episode(id, snumber, enumber)
	if err != nil {
		l.Errorf(err.Error())
		http.Error(w, "HTTP 500 : Internal Server Error", 500)
		return
	}
	context := struct {
		Title   string
		TV      tmdb.TV
		Season  tmdb.Season
		Episode tmdb.Episode
	}{
		"tvt.io",
		tv,
		season,
		episode,
	}
	t := template.Must(template.ParseFiles(
		"templates/episode.html",
		"templates/partials/facebook.html",
		"templates/partials/footer.html",
		"templates/partials/javascript.html",
		"templates/partials/css.html",
	))
	t.Execute(w, context)
}
