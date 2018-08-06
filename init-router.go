package main

import (
	"net/http"
	"github.com/gorilla/mux"
	. "golang-todolist/controllers"
)

func InitRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexController.Index).Name("index")
	r.HandleFunc("/test", TestController.Test).Name("test")
	return r
}
