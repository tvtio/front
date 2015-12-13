package routes

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/tvtio/front/logger"
)

// Terms is the /terms route
func Terms(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles(
		"templates/terms.html",
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
	}{
		"tvt.io",
	}
	t.Execute(w, context)
}
