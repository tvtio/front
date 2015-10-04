package routes

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/tvtio/front/catalog"
	"github.com/tvtio/front/models"
	"github.com/tvtio/front/tmdb"
)

// Movie is the /movie/:id route
func Movie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	session := sessions.GetSession(r)

	// Get user
	userid := fmt.Sprint(session.Get("user"))
	user, err := models.GetUser(userid)
	if err != nil {
		log.Fatal(err)
	}

	id := ps.ByName("id")
	movie, err := catalog.Movie(id)
	if err != nil {
		log.Fatal(err)
	}
	context := struct {
		Title string
		Movie tmdb.Movie
		User  *models.User
	}{
		"tvt.io",
		movie,
		user,
	}
	t, err := template.ParseFiles("templates/movie.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, context)
}
