package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"golang-todolist/frame" // ControllerMgr
	_ "golang-todolist/controllers" // init all controllers
)

func InitRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", frame.Dispatch("IndexController", "Index")).Name("index")
	// r.HandleFunc("/todolist/{id:[0-9]+}")
	r.HandleFunc("/test", frame.Dispatch("TestController", "Test")).Name("test")
	return r
}
