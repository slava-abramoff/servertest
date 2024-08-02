package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	router.ServeFiles("/public/*filepath", http.Dir("./public"))

	log.Fatal(http.ListenAndServe(":8080", router))

}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	http.ServeFile(w, r, "./public/index.html")
}
