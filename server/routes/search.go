package routes

import (
	"net/http"
	"net/url"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/repejota/logger"
	"github.com/tvtio/front/catalog"
	"github.com/tvtio/front/tmdb"
)

// Search is the /search?q=matrix route
func Search(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	l := logger.New("default")

	query := r.URL.Query().Get("q")
	results, err := catalog.SearchMulti(url.QueryEscape(query))
	if err != nil {
		l.Errorf(err.Error())
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
