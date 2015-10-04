package routes

import (
	"net/http"

	"github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
)

// Logout is the / route
func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	session := sessions.GetSession(r)
	session.Delete("user")
	http.Redirect(w, r, "http://localhost:8080/", http.StatusFound)
}
