package routes

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/tvtio/front/catalog"
	"github.com/tvtio/front/logger"
	"github.com/tvtio/front/models"
	"github.com/tvtio/front/tmdb"
)

// Episode is the /tv/:id/season/:sid/episode/:enumber route
func Episode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	session := sessions.GetSession(r)

	// Get user
	userid := fmt.Sprint(session.Get("user"))
	user, err := models.GetUser(userid)
	if err != nil {
		logger.Fatal(err.Error())
	}

	id := ps.ByName("id")
	tv, err := catalog.TV(id)
	if err != nil {
		logger.Fatal(err.Error())
	}
	snumber := ps.ByName("snumber")
	season, err := catalog.Season(id, snumber)
	if err != nil {
		logger.Fatal(err.Error())
	}
	enumber := ps.ByName("enumber")
	episode, err := catalog.Episode(id, snumber, enumber)
	if err != nil {
		logger.Fatal(err.Error())
	}
	context := struct {
		Title   string
		TV      tmdb.TV
		Season  tmdb.Season
		Episode tmdb.Episode
		User    *models.User
	}{
		"tvt.io",
		tv,
		season,
		episode,
		user,
	}
	t, err := template.ParseFiles(
		"templates/episode.html",
		"templates/partials/facebook.html",
		"templates/partials/footer.html",
		"templates/partials/javascript.html",
		"templates/partials/css.html",
	)
	if err != nil {
		logger.Fatal(err.Error())
	}
	t.Execute(w, context)
}
