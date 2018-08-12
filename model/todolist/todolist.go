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
	Id int `db:"id"`
	Name string `db:"name"`
}
