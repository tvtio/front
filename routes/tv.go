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

// TV is the /tv/:id route
func TV(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	tv, err := tmdb.GetTV(id)
	if err != nil {
		log.Fatal(err)
	}
	context := struct {
		Title string
		TV    tmdb.TV
		User  *models.User
	}{
		"tvt.io",
		tv,
		&user,
	}
	t, err := template.ParseFiles("templates/tv.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, context)
}
