package controllers

import (
	"net/http"
	"golang-todolist/templates"
)

func Test(w http.ResponseWriter, r *http.Request) {
	templates.TestView.Execute(w, nil)
}
