package controllers

import (
	"golang-todolist/templates"
	"golang-todolist/frame"
	"golang-todolist/model/todolist"
	"log"
)

func init() {
	newController := frame.NewController("IndexController")
	newController.Actions["Index"] = func() {
			resultSet := todolist.Collection().Find()
			var todoLists []todolist.Todolist
			err := resultSet.All(&todoLists)
			if err != nil {
				log.Fatalf("resultSet.All(): %q\n", err)
			}
			newController.Render("index", templates.IndexVars{todoLists, ""})
		}
}


