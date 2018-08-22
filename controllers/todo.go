package controllers

import (
	"golang-todolist/frame"
	"golang-todolist/model"
	"log"
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
		todoListId := this.Param("todo_list_id") // route "todo_new"
		// or editing an existing todo
		if todoListId == "" {
			id := this.Param("id") // route "todo_edit"
			todo = model.FindTodo("id", id)
			todoListId = todo.TodoListId
		} else {
			todo = &model.Todo{Id: "", Name: "", TodoListId: todoListId,}
		}

		list := model.FindTodolist("id", todoListId)

		this.Render("todo/edit",
			"Todo", todo,
			"List", list)
	}

	this.Actions["Save"] = func() {
		list := model.FindTodolist("id", this.Param("todo_list_id"))

		todo := model.Todo{
			Id: this.Param("id"),
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
			this.Error(err.Error())
			return
		}
		this.Redirect(frame.URL("todolist", "id", list.Id))
	}

	this.Actions["Delete"] = func() {
		var todo *model.Todo
		id := this.Param("id")
		rs := model.Todos().Find("id", id)
		err := rs.One(&todo)
		if (err != nil) {
			log.Fatalf("rs.One(&todo): %q\n", err)
		}
		rs.Delete()
		this.Redirect(frame.URL("todolist", "id", todo.TodoListId))
	}
}
