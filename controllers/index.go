package controllers

import (
	"net/http"
	"golang-todolist/templates"
	"golang-todolist/frame"
	"golang-todolist/model"
	"log"
)

var IndexController IndexControllerType

func init() {
	IndexController = IndexControllerType{}
}

type IndexControllerType struct {
	frame.Controller
}

func (this *IndexControllerType) Index(w http.ResponseWriter, r *http.Request) {
	resultSet := frame.DB().Collection("todo_list").Find()
	var todoLists []model.Todolist
	err := resultSet.All(&todoLists)
	if err != nil {
		log.Fatalf("res.All(): %q\n", err)
	}

	this.Render(w, "index", templates.IndexVars{todoLists, ""})
}
