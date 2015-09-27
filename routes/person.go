package routes

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/tvtio/front/tmdb"
)

// Person is the /movie/:id route
func Person(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	person, err := tmdb.GetPerson(id)
	if err != nil {
		log.Fatal(err)
	}
	context := struct {
		Title  string
		Person tmdb.Person
	}{
		"tvt.io",
		person,
	}
	t, err := template.ParseFiles("templates/person.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, context)
}
