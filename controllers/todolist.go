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
		id := controller.Param("id")
		var list *todolist.Todolist
		// update an existing list
		if (id != "") {
			rs := todolist.Collection().Find("id", id)
			err := rs.One(&list)
			if (err != nil) {
				log.Fatalf("rs.One(&list): %q\n", err)
			}
		}
		// or create a new one
		if (list == nil) {
			list = &todolist.Todolist{Id: "", Name: "",}
		}
		controller.Render("todolist/edit", list)
	}

	controller.Actions["Save"] = func() {
		list := todolist.Todolist{Id: controller.Param("id"), Name: controller.Param("name")}
		err := list.Save()
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


