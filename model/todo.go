package model

import (
	"upper.io/db.v2" // required for db.Collection
	"golang-todolist/frame"
)

// implements frame.Record
type Todo struct {
	Id string `db:"id"`
	Name string `db:"name"`
	TodoListId string `db:"todo_list_id"`
}

func Todos() db.Collection {
	return frame.DB().Collection("todo")
}

func (this *Todo) PrimaryKey() string {
	return this.Id
}

func (this *Todo) SetPrimaryKey(id string) {
	this.Id = id
}

func (this *Todo) Collection() db.Collection {
	return Todos()
}
