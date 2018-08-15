package controllers

import (
	"golang-todolist/frame"
	"golang-todolist/model/todolist"
	"log"
)

func init() {
	controller := frame.NewController("IndexController")

	controller.Actions["Index"] = func() {
		resultSet := todolist.Collection().Find()
		var todoLists []todolist.Todolist
		err := resultSet.All(&todoLists)
		if err != nil {
			log.Fatalf("resultSet.All(): %q\n", err)
		}
		// controller.Render("index", templates.IndexVars{todoLists, ""})
		controller.Render("index", map[string]interface{}{
			"Results": todoLists,
		})
	}

	controller.Actions["Create"] = func() {
	}
}


