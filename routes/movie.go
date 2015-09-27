package routes

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/tvtio/front/tmdb"
)

// Movie is the /movie/:id route
func Movie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	movie, err := tmdb.GetMovie(id)
	if err != nil {
		log.Fatal(err)
	}
	context := struct {
		Title string
		Movie tmdb.Movie
	}{
		"tvt.io",
		movie,
	}
	t, err := template.ParseFiles("templates/movie.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, context)
}
