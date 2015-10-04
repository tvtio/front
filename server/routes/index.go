package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/tvtio/front/catalog"
	"github.com/tvtio/front/models"
	"github.com/tvtio/front/tmdb"
)

// Index is the / route
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session := sessions.GetSession(r)

	var user models.User
	userbytes := session.Get("user")
	if userbytes != nil {
		f := userbytes.([]byte)
		err := json.Unmarshal(f, &user)
		if err != nil {
			log.Fatal(err)
		}
	}

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
		User          *models.User
	}{
		"tvt.io",
		popularMovies,
		popularTV,
		&user,
	}
	t.Execute(w, context)
}
