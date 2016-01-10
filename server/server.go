// Copyright 2016 The front Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/julienschmidt/httprouter"
	"github.com/repejota/logger"
	"github.com/tvtio/front/server/routes"
)

// Start starts the HTTP Server of the application
// * Define the routes
// * Setup middleware engine
func Start() error {
	l := logger.New("default")

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
	router.GET("/tv/:id/season/:snumber", routes.Season)
	router.GET("/tv/:id/season/:snumber/episode/:enumber", routes.Episode)

	router.GET("/person/:id", routes.Person)

	router.GET("/about", routes.About)
	router.GET("/policy", routes.Policy)
	router.GET("/terms", routes.Terms)
	router.GET("/support", routes.Support)

	router.ServeFiles("/css/*filepath", http.Dir("/go/src/github.com/tvtio/front/templates/css"))
	router.ServeFiles("/js/*filepath", http.Dir("/go/src/github.com/tvtio/front/templates/js"))
	router.ServeFiles("/img/*filepath", http.Dir("/go/src/github.com/tvtio/front/templates/img"))

	// Apply middleware to the router
	middle.UseHandler(router)

	l.Infof("Listening at :8080")

	// Start server
	return http.ListenAndServe(":8080", middle)
}
