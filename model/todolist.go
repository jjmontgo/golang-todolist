package model

import (
	"log"
	"upper.io/db.v2" // required for db.Collection
	"golang-todolist/frame"
)

// implements frame.Record
type Todolist struct {
	Id string `db:"id"`
	Name string `db:"name"`
}

func Todolists() db.Collection {
	return frame.DB().Collection("todo_list")
}

func (this *Todolist) PrimaryKey() string {
	return this.Id
}

func (this *Todolist) SetPrimaryKey(id string) {
	this.Id = id
}

func (this *Todolist) Collection() db.Collection {
	return Todolists()
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
