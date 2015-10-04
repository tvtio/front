package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"text/template"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/tvtio/front/catalog"
	"github.com/tvtio/front/models"
	"github.com/tvtio/front/tmdb"
)

// Search is the /search?q=matrix route
func Search(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	query := r.URL.Query().Get("q")
	results, err := catalog.SearchMovies(url.QueryEscape(query))
	if err != nil {
		log.Fatal(err)
	}
	context := struct {
		Title   string
		Query   string
		Results tmdb.SearchMovieResult
		User    *models.User
	}{
		"tvt.io",
		query,
		results,
		&user,
	}
	t, err := template.ParseFiles("templates/search.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, context)
}
