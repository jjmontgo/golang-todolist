package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"golang-todolist/frame" // Dispatch(), Registry
)

func InitRouter() http.Handler {
	r := mux.NewRouter()
	frame.Registry.Router = r
	r.HandleFunc("/", frame.Dispatch("IndexController", "Index")).Name("index")
	r.HandleFunc("/todolist/create", frame.Dispatch("IndexController", "Create")).Name("todolist_create")
	// r.HandleFunc("/todolist/{id:[0-9]+}")
	r.HandleFunc("/test", frame.Dispatch("TestController", "Test")).Name("test")
	return r
}
