package main

import (
	"compress/gzip"
	"fmt"
	"log"
	"net/http"

	"github.com/carbocation/interpose"
	"github.com/carbocation/interpose/middleware"
	"github.com/julienschmidt/httprouter"
	"github.com/tvtio/front/routes"
)

func main() {
	host := "localhost"
	port := "8080"
	serverAddress := host + ":" + port

	middle := interpose.New()
	// Using Gorilla framework's combined logger
	middle.Use(middleware.GorillaLog())
	//Using Negroni's Gzip functionality
	middle.Use(middleware.NegroniGzip(gzip.DefaultCompression))

	router := httprouter.New()
	router.GET("/", routes.Index)
	router.GET("/search", routes.Search)
	router.GET("/movie/:id", routes.Movie)
	router.GET("/tv/:id", routes.TV)
	router.GET("/person/:id", routes.Person)
	router.GET("/about", routes.About)
	router.GET("/login", routes.Login)
	router.ServeFiles("/css/*filepath", http.Dir("./templates/css"))
	router.ServeFiles("/js/*filepath", http.Dir("./templates/js"))
	router.ServeFiles("/img/*filepath", http.Dir("./templates/img"))

	middle.UseHandler(router)

	fmt.Println("Listening at http://" + serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, middle))
}
