package server

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/julienschmidt/httprouter"
	"github.com/tvtio/front/config"
	"github.com/tvtio/front/routes"
)

// Start ...
func Start(configuration config.Configuration) error {
	// Setup middleware
	middle := negroni.Classic()

	store := cookiestore.New([]byte("keyboard cat"))
	middle.Use(sessions.Sessions("tvtio", store))

	// Setup router
	router := httprouter.New()
	router.GET("/", routes.Index)
	router.GET("/search", routes.Search)
	router.GET("/movie/:id", routes.Movie)
	router.GET("/tv/:id", routes.TV)
	router.GET("/person/:id", routes.Person)
	router.GET("/about", routes.About)
	router.GET("/policy", routes.Policy)
	router.GET("/terms", routes.Terms)
	router.GET("/support", routes.Support)
	router.GET("/login", routes.Login)
	router.GET("/logout", routes.Logout)
	router.GET("/auth/twitter", routes.AuthTwitter)
	router.GET("/auth/facebook", routes.AuthFacebook)
	router.GET("/auth/facebook/callback", routes.AuthFacebookCallback)
	router.ServeFiles("/css/*filepath", http.Dir(configuration.Templates+"/css"))
	router.ServeFiles("/js/*filepath", http.Dir(configuration.Templates+"/js"))
	router.ServeFiles("/img/*filepath", http.Dir(configuration.Templates+"/img"))

	middle.UseHandler(router)

	// Start server
	log.Println("Listening at " + configuration.ServerURL())
	return http.ListenAndServe(configuration.ServerAddress(), middle)
}
