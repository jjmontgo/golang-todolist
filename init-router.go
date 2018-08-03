package main

import (
	"net/http"
	"github.com/gorilla/mux"
	. "golang-todolist/controllers"
)

func InitRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", Index).Name("index")
	r.HandleFunc("/test", Test).Name("test")
	return r
}
