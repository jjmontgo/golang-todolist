package model

// implements frame.Record
type Todo struct {
	Id uint `sql:"type:int PRIMARY KEY"`
	Name string `sql:"type:varchar(250)"`
	TodoListId uint `sql:"type:int"`
}
