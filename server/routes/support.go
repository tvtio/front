package routes

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// Support is the /policy route
func Support(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := template.Must(template.ParseFiles(
		"templates/support.html",
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
