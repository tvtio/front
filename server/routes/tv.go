package routes

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/mgutz/ansi"
	"github.com/tvtio/front/catalog"
	"github.com/tvtio/front/models"
	"github.com/tvtio/front/tmdb"
)

// TV is the /tv/:id route
func TV(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	session := sessions.GetSession(r)

	// Get user
	userid := fmt.Sprint(session.Get("user"))
	user, err := models.GetUser(userid)
	if err != nil {
		log.Fatal(ansi.Color("FATAL: ", "red"), err)
	}

	id := ps.ByName("id")
	tv, err := catalog.TV(id)
	if err != nil {
		log.Fatal(ansi.Color("FATAL: ", "red"), err)
	}
	context := struct {
		Title string
		TV    tmdb.TV
		User  *models.User
	}{
		"tvt.io",
		tv,
		user,
	}
	t, err := template.ParseFiles(
		"templates/tv.html",
		"templates/partials/facebook.html",
		"templates/partials/footer.html",
		"templates/partials/javascript.html",
		"templates/partials/css.html",
	)
	if err != nil {
		log.Fatal(ansi.Color("FATAL: ", "red"), err)
	}
	t.Execute(w, context)
}
