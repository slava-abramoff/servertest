package main

import (
	"log"
	"net/http"
)

func main() {
	OpenDB()
	SetupRouter()
	log.Println("Server is starting...")
	log.Fatal(http.ListenAndServe(":8080", Router))
}
