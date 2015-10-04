package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/tvtio/front/models"
)

// Terms is the /policy route
func Terms(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	t, err := template.ParseFiles("templates/terms.html")
	if err != nil {
		log.Fatal(err)
	}
	context := struct {
		Title string
		User  *models.User
	}{
		"tvt.io",
		&user,
	}
	t.Execute(w, context)
}
