package model

// implements frame.Record
type Todo struct {
	Id uint `sql:"type:int PRIMARY KEY" json:"id"`
	Name string `sql:"type:varchar(250)" json:"name"`
	TodoListId uint `sql:"type:int" json:"todo_list_id"`
}
