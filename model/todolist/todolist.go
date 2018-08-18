package todolist

import (
	"golang-todolist/frame"
	"upper.io/db.v2" // db package
)

const TableName = `todo_list`

func Collection() db.Collection {
	return frame.DB().Collection(TableName)
}

type Todolist struct {
	Id string `db:"id"`
	Name string `db:"name"`
}

func (this *Todolist) Save() error {
	var err error
	var id interface{}
	if (this.Id == "") {
		id, err = Collection().Insert(this)
		this.Id = frame.ToString(id)
	} else {
		err = Collection().Find("id", this.Id).Update(this)
	}

	return err
}
