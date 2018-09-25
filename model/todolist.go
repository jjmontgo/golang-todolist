package model

type TodoList struct {
	Id uint `sql:"type:int PRIMARY KEY"`
	Name string `sql:"type:varchar(250)"`
	Todos []Todo
	MediaAttachment *MediaAttachment `gorm:"polymorphic:Ref;"`
}
