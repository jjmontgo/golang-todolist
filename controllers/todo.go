package controllers

import (
	"golang-todolist/frame"
	"golang-todolist/model"
	// "log"
	"strings"
)

func init() {
	this := frame.NewController("Todo")

	this.IsAccessible = func(actionName string) bool {
		return this.UserIsLoggedIn()
	}

	this.Actions["Index"] = func() {
		var todoList model.TodoList
		this.DB().First(&todoList, this.ParamUint("id"))
		var todos []model.Todo
		this.DB().Model(&todoList).Related(&todos)

		this.RenderJSON(
			"todolist", todoList,
			"todos", todos,
		)
	}

	this.Actions["Edit"] = func() {
		var todo model.Todo
		// creating a new todo
		todoListId := this.ParamUint("todo_list_id")
		// or editing an existing todo
		if todoListId == 0 {
			id := this.ParamUint("id")
			this.DB().First(&todo, id)
			todoListId = todo.TodoListId
		} else {
			todo = model.Todo{Name: "", TodoListId: todoListId,}
		}

		var todoList model.TodoList
		this.DB().First(&todoList, todoListId)

		this.RenderJSON("todo", todo)
	}

	this.Actions["Save"] = func() {
		var todoList model.TodoList
		todoListId := this.ParamUint("todo_list_id")
		this.DB().First(&todoList, todoListId)

		todo := model.Todo{
			Id: this.ParamUint("id"),
			Name: strings.Trim(this.Param("name"), " "),
			TodoListId: todoList.Id,
		}

		if todo.Name == "" {
			this.RenderJsonError("message", "You must enter a name.")
			return
		}

		this.DB().Save(&todo)
	}

	this.Actions["Delete"] = func() {
		id := this.ParamUint("id")
		var todo model.Todo
		this.DB().First(&todo, id)
		this.DB().Delete(&todo)
	}
}
