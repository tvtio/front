// Copyright 2016 The front Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"net/url"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/repejota/logger"
	"github.com/tvtio/front/catalog"
	"github.com/tvtio/tmdb"
)

// Search is the /search?q=matrix route
func Search(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	l := logger.New("default")

	query := r.URL.Query().Get("q")
	results, err := catalog.SearchMulti(url.QueryEscape(query))
	if err != nil {
		l.Errorf(err.Error())
		http.Error(w, "HTTP 500 : Internal Server Error", 500)
		return
	}
	context := struct {
		Title   string
		Query   string
		Results tmdb.SearchMultiResult
	}{
		"tvt.io",
		query,
		results,
	}
	t := template.Must(template.ParseFiles(
		"templates/search.html",
		"templates/partials/facebook.html",
		"templates/partials/footer.html",
		"templates/partials/javascript.html",
		"templates/partials/css.html",
	))
	t.Execute(w, context)
}
