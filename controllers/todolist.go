package controllers


import (
	// "log"
	"golang-todolist/frame"
	"golang-todolist/model"
	"strings"
)

func init() {
	this := frame.NewController("Todolist")

	this.IsAccessible = func(actionName string) bool {
		return frame.UserIsLoggedIn()
	}

	this.Actions["Index"] = func() {
		todoLists := model.FindTodolists()
		this.Render("todolist/index", "Results", todoLists)
	}

	this.Actions["Edit"] = func() {
		id := frame.StringToUint(this.Param("id"))
		var list *model.Todolist
		// update an existing list
		if id != 0 {
			list = model.FindTodolist("id", id)
		}
		// or create a new one
		if list == nil {
			list = &model.Todolist{Name: "",}
		}
		this.Render("todolist/edit", "List", list)
	}

	this.Actions["Save"] = func() {
		list := model.Todolist{Id: frame.StringToUint(this.Param("id")), Name: this.Param("name")}
		list.Name = strings.Trim(this.Param("name"), " ")
		if list.Name == "" {
			this.Render("todolist/edit",
				"List", list,
				"Error", "You must enter a name.")
			return
		}

		err := frame.SaveRecord(&list)
		if err != nil {
			this.Error(err)
			return
		}
		this.Redirect(frame.URL("index"))
	}

	this.Actions["Delete"] = func() {
		id := frame.StringToUint(this.Param("id"))
		todolist := model.FindTodolist("id", id)
		todolist.Delete()
		this.Redirect(frame.URL("index"))
	}
}


