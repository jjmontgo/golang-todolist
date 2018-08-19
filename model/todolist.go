package model

import (
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
