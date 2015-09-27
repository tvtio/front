package routes

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/tvtio/front/tmdb"
)

// TV is the /tv/:id route
func TV(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	tv, err := tmdb.GetTV(id)
	if err != nil {
		log.Fatal(err)
	}
	context := struct {
		Title string
		TV    tmdb.TV
	}{
		"tvt.io",
		tv,
	}
	t, err := template.ParseFiles("templates/tv.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, context)
}
