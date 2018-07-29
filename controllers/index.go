package controllers

import (
	"net/http"
	"golang-todolist/templates"
)

func Index(w http.ResponseWriter, r *http.Request) {
	templates.LoadIndexTemplate().Execute(w, nil)
}
