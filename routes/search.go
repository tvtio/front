package routes

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/tvtio/front/tmdb"
)

// Search is the /search?q=matrix route
func Search(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := r.URL.Query().Get("q")
	results, err := tmdb.SearchMovie(query)
	if err != nil {
		log.Fatal(err)
	}
	context := struct {
		Title   string
		Query   string
		Results tmdb.SearchMovieResult
	}{
		"tvt.io",
		query,
		results,
	}
	t, err := template.ParseFiles("templates/search.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, context)
}
