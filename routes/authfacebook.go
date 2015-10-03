package routes

import (
	"log"
	"net/http"
	"text/template"

	"golang.org/x/oauth2"

	"github.com/julienschmidt/httprouter"
)

// AuthFacebook is the /auth/twitter route
func AuthFacebook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	oauthConfig := &oauth2.Config{ //setup
		ClientID:     "910962655664722",
		ClientSecret: "cdd845180e7223f34f4e2617360414cf",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/dialog/oauth",
			TokenURL: "https://graph.facebook.com/oauth/access_token",
		},
		RedirectURL: "http://localhost:8080",
	}

	url := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusFound)

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
