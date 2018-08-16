package controllers

import (
	// "fmt"
	"log"
	"golang-todolist/frame"
	"golang-todolist/model/todolist"
)

func init() {
	controller := frame.NewController("Todolist")

	controller.Actions["Index"] = func() {
		resultSet := todolist.Collection().Find()
		var todoLists []todolist.Todolist
		err := resultSet.All(&todoLists)
		if err != nil {
			log.Fatalf("resultSet.All(): %q\n", err)
		}

		controller.Render("todolist/index", map[string]interface{}{
			"Results": todoLists,
		})
	}

	controller.Actions["Edit"] = func() {
		controller.Render("todolist/edit", nil)
	}

	controller.Actions["Save"] = func() {
		name := controller.Param("name")
		_, err := todolist.Collection().Insert(todolist.Todolist{Name: name})
		if err != nil {
			controller.Error(err.Error())
			return
		}
		controller.Redirect(controller.Route("index"))
	}

	controller.Actions["Delete"] = func() {
		id := controller.Param("id")
		todolist.Collection().Find("id", id).Delete()
		controller.Redirect(controller.Route("index"))
	}
}


