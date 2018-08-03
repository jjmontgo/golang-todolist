package controllers

import (
	"net/http"
	"golang-todolist/templates"
	"golang-todolist/model"
)

func Index(w http.ResponseWriter, r *http.Request) {
	db := model.DB()
	results := []templates.Todolist{}
	rows, err := db.Query("SELECT * FROM todo_list")
	if (err != nil) {
		panic("Failed to load todo lists")
	}
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		list := templates.Todolist{id, name}
		results = append(results, list)
	}
	templates.ExecuteIndexTemplate(w, templates.TplIndexVars{results, ""})
}
