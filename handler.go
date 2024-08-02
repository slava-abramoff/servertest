package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Функция отправки html-документа index.html
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/index.html")
	log.Println("Somebody connected to server")
}

// Функция авторизации клиента через данные формы. При успешной авторизации перенаправляет пользователя, иначе выдаёт ошибку.
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

// Функция отправки html-документа home.html
func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "./public/home.html")
}

// Функция отправки пользователю файла Могильщик
func Download(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "public/storage/13. Могильщик.mp3")
	log.Println("Somebody downloaded chanson")
}
