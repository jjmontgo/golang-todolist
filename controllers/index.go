package controllers

import (
	"net/http"
	"golang-todolist/templates"
	"golang-todolist/model"
)

type todolist struct {
	Id int
	Name string
}

type templateVars struct {
	Results []todolist
	IndexURL string
}

func Index(w http.ResponseWriter, r *http.Request) {
	db := model.DB()
	results := []todolist{}
	rows, err := db.Query("SELECT * FROM todo_list")
	if (err != nil) {
		panic("Failed to load todo lists")
	}
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		list := todolist{id, name}
		results = append(results, list)
	}

	templates.IndexView.Execute(w, templateVars{results, ""})
}
