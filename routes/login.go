package routes

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// Login is the /login route
func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles("templates/login.html")
	if err != nil {
		log.Fatal(err)
	}
	context := struct {
		Title string
	}{
		"tvt.io",
	}
	t.Execute(w, context)
}
