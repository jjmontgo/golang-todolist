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
		db := frame.GORM()
		var todoList model.TodoList
		db.First(&todoList, frame.StringToUint(this.Param("id")))
		var todos []model.Todo
		db.Model(&todoList).Related(&todos)

		this.Render(
			"todo/index",
			"List", todoList,
			"Todos", todos,
		)
	}

	this.Actions["Edit"] = func() {
		db := frame.GORM()
		var todo model.Todo
		// creating a new todo
		todoListId := frame.StringToUint(this.Param("todo_list_id"))
		// or editing an existing todo
		if todoListId == 0 {
			id := frame.StringToUint(this.Param("id"))
			db.First(&todo, id)
			todoListId = todo.TodoListId
		} else {
			todo = model.Todo{Name: "", TodoListId: todoListId,}
		}

		var todoList model.TodoList
		db.First(&todoList, todoListId)

		this.Render("todo/edit",
			"Todo", todo,
			"List", todoList)
	}

	this.Actions["Save"] = func() {
		db := frame.GORM()

		var todoList model.TodoList
		todoListId := frame.StringToUint(this.Param("todo_list_id"))
		db.First(&todoList, todoListId)

		todo := model.Todo{
			Id: frame.StringToUint(this.Param("id")),
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

		db.Save(&todo)
		this.Redirect(frame.URL("todolist", "id", frame.UintToString(todoList.Id)))
	}

	this.Actions["Delete"] = func() {
		db := frame.GORM()
		id := frame.StringToUint(this.Param("id"))
		var todo model.Todo
		db.First(&todo, id)
		db.Delete(&todo)
		this.Redirect(frame.URL("todolist", "id", frame.UintToString(todo.TodoListId)))
	}
}
