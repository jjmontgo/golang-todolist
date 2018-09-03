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
		list := model.FindTodolist("id", this.Param("id"))
		todos := list.GetTodos()
		this.Render(
			"todo/index",
			"List", list,
			"Todos", todos,
		)
	}

	this.Actions["Edit"] = func() {
		var todo *model.Todo
		// creating a new todo
		todoListId := frame.StringToUint(this.Param("todo_list_id"))
		// or editing an existing todo
		if todoListId == 0 {
			todo = model.FindTodo("id", this.Param("id"))
			todoListId = todo.TodoListId
		} else {
			todo = &model.Todo{Name: "", TodoListId: todoListId,}
		}

		list := model.FindTodolist("id", todoListId)

		this.Render("todo/edit",
			"Todo", todo,
			"List", list)
	}

	this.Actions["Save"] = func() {
		list := model.FindTodolist("id", this.Param("todo_list_id"))

		todo := model.Todo{
			Id: frame.StringToUint(this.Param("id")),
			Name: strings.Trim(this.Param("name"), " "),
			TodoListId: list.Id,
		}

		if todo.Name == "" {
			this.Render("todo/edit",
				"Todo", todo,
				"List", list,
				"Error", "You must enter a name.")
			return
		}

		err := frame.SaveRecord(&todo)
		if err != nil {
			this.Error(err)
			return
		}
		this.Redirect(frame.URL("todolist", "id", frame.UintToString(list.Id)))
	}

	this.Actions["Delete"] = func() {
		todo := model.FindTodo("id", this.Param("id"))
		frame.DeleteRecord(todo)
		this.Redirect(frame.URL("todolist", "id", frame.UintToString(todo.TodoListId)))
	}
}
