package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tvtio/front/routes"
)

func main() {
	host := "localhost"
	port := "8080"
	serverAddress := host + ":" + port
	router := httprouter.New()
	router.GET("/", routes.Index)
	router.GET("/search", routes.Search)
	router.GET("/about", routes.About)
	router.ServeFiles("/css/*filepath", http.Dir("./templates/css"))
	router.ServeFiles("/js/*filepath", http.Dir("./templates/js"))
	router.ServeFiles("/img/*filepath", http.Dir("./templates/img"))
	fmt.Println("Listening at http://" + serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, router))
}
