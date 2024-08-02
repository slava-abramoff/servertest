package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var Router *httprouter.Router

func SetupRouter() {
	Router = httprouter.New()
	Router.GET("/", Index)
	Router.GET("/home", Home)
	Router.GET("/home/download", Download)
	Router.POST("/auth", Auth)

	Router.ServeFiles("/public/*filepath", http.Dir("./public"))
}
