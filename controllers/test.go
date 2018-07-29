package controllers

import (
	"net/http"
	"golang-todolist/templates"
)

func Test(w http.ResponseWriter, r *http.Request) {
	templates.LoadTestTemplate().Execute(w, nil)
}
