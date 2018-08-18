package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"golang-todolist/frame" // Dispatch(), Registry
)

func InitRouter() http.Handler {
	r := mux.NewRouter()
	frame.Registry.Router = r

	r.HandleFunc("/", frame.Dispatch("Todolist", "Index")).Name("index")

	r.HandleFunc("/todolist/new", frame.Dispatch("Todolist", "Edit")).Name("todolist_new")

	r.HandleFunc("/todolist/edit/{id:[0-9]+}", frame.Dispatch("Todolist", "Edit")).Name("todolist_edit")

	r.HandleFunc("/todolist/save", frame.Dispatch("Todolist", "Save")).
		Methods("POST").
		Name("todolist_save")

	r.HandleFunc("/todolist/delete/{id}", frame.Dispatch("Todolist", "Delete")).
		Methods("POST").
		Name("todolist_delete")

	r.HandleFunc("/test", frame.Dispatch("TestController", "Test")).Name("test")

	return r
}
