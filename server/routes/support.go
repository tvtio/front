package routes

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/tvtio/front/models"
)

// Support is the /policy route
func Support(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session := sessions.GetSession(r)

	// Get user
	userid := fmt.Sprint(session.Get("user"))
	user, err := models.GetUser(userid)
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles(
		"templates/support.html",
		"templates/partials/facebook.html",
		"templates/partials/javascript.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	context := struct {
		Title string
		User  *models.User
	}{
		"tvt.io",
		user,
	}
	t.Execute(w, context)
}
