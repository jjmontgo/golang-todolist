package controllers

import (
	"golang-todolist/frame"
	"golang-todolist/model"
	"log"
)

func init() {
	this := frame.NewController("Todo")

	this.Actions["Index"] = func() {
		rs := model.Todolists().Find("id", this.Param("id"))
		var list *model.Todolist
		err := rs.One(&list)
		if (err != nil) {
			log.Fatalf("rs.One(&list): %q\n", err)
		}
		this.Render("todo/index", "List", list)
	}
}
