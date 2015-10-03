package main

import (
	"compress/gzip"
	"fmt"
	"log"
	"net/http"

	"github.com/carbocation/interpose"
	"github.com/carbocation/interpose/middleware"
	"github.com/julienschmidt/httprouter"
	"github.com/tvtio/front/config"
	"github.com/tvtio/front/routes"
)

func main() {
	// Load configuration
	// TODO: How to share this config with the handlers?
	configuration, err := config.Load("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	// Setup middleware
	middle := interpose.New()
	// Using Gorilla framework's combined logger
	// middle.Use(middleware.GorillaLog())
	// Using Negroni's Gzip functionality
	middle.Use(middleware.NegroniGzip(gzip.DefaultCompression))

	// Setup router
	router := httprouter.New()
	router.GET("/", routes.Index)
	router.GET("/search", routes.Search)
	router.GET("/movie/:id", routes.Movie)
	router.GET("/tv/:id", routes.TV)
	router.GET("/person/:id", routes.Person)
	router.GET("/about", routes.About)
	router.GET("/login", routes.Login)
	router.GET("/auth/twitter", routes.AuthTwitter)
	router.GET("/auth/facebook", routes.AuthFacebook)
	router.GET("/auth/facebook/callback", routes.AuthFacebookCallback)
	router.ServeFiles("/css/*filepath", http.Dir(configuration.Templates+"/css"))
	router.ServeFiles("/js/*filepath", http.Dir(configuration.Templates+"/js"))
	router.ServeFiles("/img/*filepath", http.Dir(configuration.Templates+"/img"))

	middle.UseHandler(router)

	// Start server
	fmt.Println("Listening at " + configuration.ServerURL())
	log.Fatal(http.ListenAndServe(configuration.ServerAddress(), middle))
}
