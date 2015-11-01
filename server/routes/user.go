package routes

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/tvtio/front/logger"
	"github.com/tvtio/front/models"
)

// UserMe is the /user/me route
func UserMe(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session := sessions.GetSession(r)

	// Get user
	userid := fmt.Sprint(session.Get("user"))
	user, err := models.GetUser(userid)
	if err != nil {
		logger.Fatal(err.Error())
	}

	t, err := template.ParseFiles(
		"templates/user.me.html",
		"templates/partials/facebook.html",
		"templates/partials/footer.html",
		"templates/partials/javascript.html",
		"templates/partials/css.html",
	)
	if err != nil {
		logger.Fatal(err.Error())
	}
	context := struct {
		Title string
		User  *models.User
	}{
		"tvt.io - My Profile",
		user,
	}
	t.Execute(w, context)
}
