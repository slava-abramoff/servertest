package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var Router *httprouter.Router

// Функция инициализации путей httprouter
func SetupRouter() {
	Router = httprouter.New()
	Router.GET("/", Index)
	Router.GET("/home", Home)
	Router.GET("/home/download", Download)
	Router.POST("/auth", Auth)

	// Позволяет серверу обрабатывать статические файлы по указанному пути.
	Router.ServeFiles("/public/*filepath", http.Dir("./public"))
}
