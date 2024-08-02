package main

import (
	"log"
	"net/http"
)

func main() {
	SetupRouter()
	log.Println("Server is starting...")
	log.Fatal(http.ListenAndServe(":8080", Router))
}
