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
	"github.com/tvtio/front/tmdb"
)

// Season is the /tv/:id/season/:sid route
func Season(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	l := logger.New("default")

	id := ps.ByName("id")
	tv, err := catalog.TV(id)
	if err != nil {
		l.Errorf(err.Error())
	}
	snumber := ps.ByName("snumber")
	season, err := catalog.Season(id, snumber)
	if err != nil {
		l.Errorf(err.Error())
	}
	context := struct {
		Title  string
		TV     tmdb.TV
		Season tmdb.Season
	}{
		"tvt.io",
		tv,
		season,
	}
	t := template.Must(template.ParseFiles(
		"templates/season.html",
		"templates/partials/facebook.html",
		"templates/partials/footer.html",
		"templates/partials/javascript.html",
		"templates/partials/css.html",
	))
	t.Execute(w, context)
}
