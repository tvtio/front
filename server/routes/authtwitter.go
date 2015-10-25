package routes

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"golang.org/x/oauth2"

	"github.com/julienschmidt/httprouter"
)

// AuthTwitter is the /auth/twitter route
func AuthTwitter(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	oauthConfig := &oauth2.Config{
		ClientID:     "P1LBfEJNOHDaTSOQrmAb5Gvil",
		ClientSecret: "8B6wkuXDloxrmx1uUZlQEmaSnUzVU3n48XYEcHfzFe33sNu8KK",

		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.twitter.com/oauth/authorize",
			TokenURL: "https://api.twitter.com/oauth/token",
		},
	}

	url := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOnline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)
	http.Redirect(w, r, url, http.StatusFound)

	t, err := template.ParseFiles(
		"templates/login.html",
		"templates/partials/facebook.html",
		"templates/partials/footer.html",
		"templates/partials/javascript.html",
		"templates/partials/css.html",
	)
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
