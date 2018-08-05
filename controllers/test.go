package controllers

import (
	"net/http"
	"golang-todolist/frame"
)

func Test(w http.ResponseWriter, r *http.Request) {
	frame.ViewMgr.Get("test").Render(w, nil)
}
