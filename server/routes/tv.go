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

// TV is the /tv/:id route
func TV(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	l := logger.New("default")

	id := ps.ByName("id")
	tv, err := catalog.TV(id)
	if err != nil {
		l.Errorf(err.Error())
	}
	context := struct {
		Title string
		TV    tmdb.TV
	}{
		"tvt.io",
		tv,
	}
	t := template.Must(template.ParseFiles(
		"templates/tv.html",
		"templates/partials/facebook.html",
		"templates/partials/footer.html",
		"templates/partials/javascript.html",
		"templates/partials/css.html",
	))
	t.Execute(w, context)
}
