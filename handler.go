package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/index.html")
}

func Auth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "admin" && password == "admin" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		log.Printf("User %s login successfully", username)
	} else {
		http.Error(w, "Invalid login or password", http.StatusUnauthorized)
		log.Printf("User %s failed to login", username)
	}
}

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/home.html")
}

func Download(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "public/13. Могильщик.mp3")
}
