package controllers

import "strings"

import (
	// "fmt"
	"log"
	"golang-todolist/frame"
	"golang-todolist/model"
)

func init() {
	this := frame.NewController("Todolist")

	this.Actions["Index"] = func() {
		resultSet := model.Todolists().Find()
		var todoLists []model.Todolist
		err := resultSet.All(&todoLists)
		if err != nil {
			log.Fatalf("resultSet.All(): %q\n", err)
		}

		this.Render("todolist/index", "Results", todoLists)
	}

	this.Actions["Edit"] = func() {
		id := this.Param("id")
		var list *model.Todolist
		// update an existing list
		if (id != "") {
			rs := model.Todolists().Find("id", id)
			err := rs.One(&list)
			if (err != nil) {
				log.Fatalf("rs.One(&list): %q\n", err)
			}
		}
		// or create a new one
		if (list == nil) {
			list = &model.Todolist{Id: "", Name: "",}
		}
		this.Render("todolist/edit", "List", list)
	}

	this.Actions["Save"] = func() {
		list := model.Todolist{Id: this.Param("id"), Name: this.Param("name")}
		list.Name = strings.Trim(this.Param("name"), " ")
		if list.Name == "" {
			this.Render("todolist/edit",
				"List", list,
				"Error", "You must enter a name.")
			return
		}

		err := list.Save()
		if err != nil {
			this.Error(err.Error())
			return
		}
		this.Redirect(frame.URL("index"))
	}

	this.Actions["Delete"] = func() {
		id := this.Param("id")
		model.Todolists().Find("id", id).Delete()
		this.Redirect(frame.URL("index"))
	}
}


