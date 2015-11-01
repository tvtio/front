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

// Index is the / route
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Get session
	session := sessions.GetSession(r)

	// Get user
	userid := fmt.Sprint(session.Get("user"))
	user, err := models.GetUser(userid)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Get popular movies
	popularMovies, err := catalog.PopularMovies()
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Get popular tv
	popularTV, err := catalog.PopularTV()
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Build template
	t, err := template.ParseFiles(
		"templates/index.html",
		"templates/partials/facebook.html",
		"templates/partials/footer.html",
		"templates/partials/javascript.html",
		"templates/partials/css.html",
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	// Template context
	context := struct {
		Title         string
		PopularMovies tmdb.SearchMovieResult
		PopularTV     tmdb.SearchTVResult
		User          *models.User
		BG1           string
	}{
		"tvt.io",
		popularMovies,
		popularTV,
		user,
		popularMovies.Results[0].BackdropPath,
	}

	// Render template
	t.Execute(w, context)
}
