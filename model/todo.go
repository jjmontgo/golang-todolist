package model

import (
	"upper.io/db.v2" // required for db.Collection
	"golang-todolist/frame"
)

func FindTodo(searchParams ...interface{}) *Todo {
	var todo *Todo
	rs := Todos().Find(searchParams...)
	err := rs.One(&todo)
	if (err != nil) {
		return nil
	}
	return todo
}

func Todos() db.Collection {
	return frame.DB().Collection("todo")
}

// implements frame.Record
type Todo struct {
	Id uint `db:"id"`
	Name string `db:"name"`
	TodoListId uint `db:"todo_list_id"`
}

// frame.Record interface
func (this *Todo) PrimaryKey() uint {
	return this.Id
}

// frame.Record interface
func (this *Todo) SetPrimaryKey(id uint) {
	this.Id = id
}

// frame.Record interface
func (this *Todo) Collection() db.Collection {
	return Todos()
}

