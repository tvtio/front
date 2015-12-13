package routes

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// Policy is the /policy route
func Policy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := template.Must(template.ParseFiles(
		"templates/policy.html",
		"templates/partials/facebook.html",
		"templates/partials/footer.html",
		"templates/partials/javascript.html",
		"templates/partials/css.html",
	))
	context := struct {
		Title string
	}{
		"tvt.io",
	}
	t.Execute(w, context)
}
