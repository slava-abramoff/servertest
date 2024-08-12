package main

import (
	"encoding/json"
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

	if FindUserFromDB(username, password) {
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

// Отправка данных о пользователях json-файлом
func GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := json.Marshal(GetUsersFromDB())
	if err != nil {
		log.Printf("Failed to marshal users %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	w.Write(data)
}

// Отправка данных о пользователе по имени json-файлом
func GetUserByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, err := GetUserByNameFromDB(ps.ByName("name"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		log.Printf("Failed to marshal user %s", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	w.Write(data)

}
