package model

import (
	"log"
	"upper.io/db.v2" // required for db.Collection
	"golang-todolist/frame"
)

func FindTodolists(searchParams ...interface{}) []Todolist {
	resultSet := Todolists().Find(searchParams...)
	var todoLists []Todolist
	err := resultSet.All(&todoLists)
	if err != nil {
		log.Fatalf("FindTodolists(): %q\n", err)
	}
	return todoLists
}

func FindTodolist(searchParams ...interface{}) *Todolist {
	var todoList *Todolist
	rs := Todolists().Find(searchParams...)
	err := rs.One(&todoList)
	if (err != nil) {
		log.Fatalf("rs.One(&list): %q\n", err)
	}
	return todoList
}

func Todolists() db.Collection {
	return frame.DB().Collection("todo_list")
}

// implements frame.Record
type Todolist struct {
	Id string `db:"id"`
	Name string `db:"name"`
}

// frame.Record interface
func (this *Todolist) PrimaryKey() string {
	return this.Id
}

// frame.Record interface
func (this *Todolist) SetPrimaryKey(id string) {
	this.Id = id
}

// frame.Record interface
func (this *Todolist) Collection() db.Collection {
	return Todolists()
}

func (this *Todolist) Delete() {
	Todos().Find("todo_list_id", this.Id).Delete()
	frame.DeleteRecord(this)
}

func (this *Todolist) GetTodos() []*Todo {
	rs := Todos().Find("todo_list_id", this.Id)
	var todos []*Todo
	err := rs.All(&todos)
	if (err != nil) {
		log.Fatalf("rs.All(&todos): %q\n", err)
	}
	return todos
}
