package model

import (
	"upper.io/db.v2" // required for db.Collection
	"golang-todolist/frame"
)

type Todolist struct {
	Id string `db:"id"`
	Name string `db:"name"`
}

func Todolists() db.Collection {
	return frame.DB().Collection("todo_list")
}

func (this *Todolist) Save() error {
	var err error
	var id interface{}
	if (this.Id == "") {
		id, err = Todolists().Insert(this)
		this.Id = frame.ToString(id)
	} else {
		err = Todolists().Find("id", this.Id).Update(this)
	}

	return err
}
