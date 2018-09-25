package controllers

import (
	"golang-todolist/frame"
	"golang-todolist/model"
	// "log"
	"strings"
)

func init() {
	this := frame.NewController("Todo")

	this.Actions["Index"] = func() {
		var todoList model.TodoList
		this.DB().First(&todoList, this.ParamUint("id"))
		var todos []model.Todo
		this.DB().Model(&todoList).Related(&todos)

		this.Render(
			"todo/index",
			"List", todoList,
			"Todos", todos,
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

		this.Render("todo/edit",
			"Todo", todo,
			"List", todoList)
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
			this.Render("todo/edit",
				"Todo", todo,
				"List", todoList,
				"Error", "You must enter a name.")
			return
		}

		this.DB().Save(&todo)
		this.Redirect(frame.URL("todolist", "id", frame.UintToString(todoList.Id)))
	}

	this.Actions["Delete"] = func() {
		id := this.ParamUint("id")
		var todo model.Todo
		this.DB().First(&todo, id)
		this.DB().Delete(&todo)
		this.Redirect(frame.URL("todolist", "id", frame.UintToString(todo.TodoListId)))
	}
}
