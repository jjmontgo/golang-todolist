package model

type Todolist struct {
	Id int `db:"id"`
	Name string `db:"name"`
}
