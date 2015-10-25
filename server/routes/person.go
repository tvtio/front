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

// Person is the /movie/:id route
func Person(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	session := sessions.GetSession(r)

	// Get user
	userid := fmt.Sprint(session.Get("user"))
	user, err := models.GetUser(userid)
	if err != nil {
		log.Fatal(err)
	}

	id := ps.ByName("id")
	person, err := catalog.Person(id)
	if err != nil {
		log.Fatal(err)
	}
	context := struct {
		Title  string
		Person tmdb.Person
		User   *models.User
	}{
		"tvt.io",
		person,
		user,
	}
	t, err := template.ParseFiles(
		"templates/person.html",
		"templates/partials/facebook.html",
		"templates/partials/footer.html",
		"templates/partials/javascript.html",
		"templates/partials/css.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, context)
}
