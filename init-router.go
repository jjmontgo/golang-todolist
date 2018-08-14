package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"golang-todolist/frame" // ControllerMgr
	_ "golang-todolist/controllers" // init all controllers
)

func InitRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", frame.Dispatch(r, "IndexController", "Index")).Name("index")
	r.HandleFunc("/todolist/create", frame.Dispatch(r, "IndexController", "Create")).Name("todolist_create")
	// r.HandleFunc("/todolist/{id:[0-9]+}")
	r.HandleFunc("/test", frame.Dispatch(r, "TestController", "Test")).Name("test")
	return r
}
