package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Search is the /search?q=matrix route
func Search(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := r.URL.Query().Get("q")

	t, _ := template.ParseFiles("templates/search.html")
	t.Execute(w, nil)
}

// Index is the / route
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func main() {
	host := "localhost"
	port := "8080"
	serverAddress := host + ":" + port

	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/search", Search)

	router.ServeFiles("/css/*filepath", http.Dir("./templates/css"))
	router.ServeFiles("/js/*filepath", http.Dir("./templates/js"))
	router.ServeFiles("/img/*filepath", http.Dir("./templates/img"))

	fmt.Println("Listening at http://" + serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, router))
}
