package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/tvtio/front/models"
	"github.com/tvtio/front/tmdb"
)

// Person is the /movie/:id route
func Person(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	id := ps.ByName("id")
	person, err := tmdb.GetPerson(id)
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
		&user,
	}
	t, err := template.ParseFiles("templates/person.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, context)
}
