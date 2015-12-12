package routes

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/mgutz/ansi"
	"github.com/tvtio/front/catalog"
	"github.com/tvtio/front/logger"
	"github.com/tvtio/front/tmdb"
)

// Movie is the /movie/:id route
func Movie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	movie, err := catalog.Movie(id)
	if err != nil {
		log.Fatal(ansi.Color("FATAL: ", "red"), err)
	}
	context := struct {
		Title string
		Movie tmdb.Movie
	}{
		"tvt.io",
		movie,
	}
	t, err := template.ParseFiles(
		"templates/movie.html",
		"templates/partials/facebook.html",
		"templates/partials/footer.html",
		"templates/partials/javascript.html",
		"templates/partials/css.html",
	)
	if err != nil {
		logger.Fatal(err.Error())
	}
	t.Execute(w, context)
}
