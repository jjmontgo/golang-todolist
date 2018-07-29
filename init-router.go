package main

import (
	"net/http"
	"github.com/gorilla/mux"
	. "golang-todolist/controllers"
)

func InitRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	r.HandleFunc("/test", Test)
	return r
}
