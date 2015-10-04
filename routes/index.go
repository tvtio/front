package routes

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/tvtio/front/catalog"
	"github.com/tvtio/front/tmdb"
)

// Index is the / route
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	popularMovies, err := catalog.PopularMovies()
	if err != nil {
		log.Fatal(err)
	}

	popularTV, err := catalog.PopularTV()
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	context := struct {
		Title         string
		PopularMovies tmdb.SearchMovieResult
		PopularTV     tmdb.SearchTVResult
	}{
		"tvt.io",
		popularMovies,
		popularTV,
	}
	t.Execute(w, context)
}
