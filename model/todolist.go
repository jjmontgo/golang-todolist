package model

import "golang-todolist/frame"

type TodoList struct {
	Id uint `sql:"type:int PRIMARY KEY"`
	Name string `sql:"type:varchar(250)"`
	Todos []Todo
	MediaAttachment *MediaAttachment `gorm:"polymorphic:Ref;"`
}

func (this *TodoList) BeforeDelete() (err error) {
	db := frame.GORM()
	db.Delete(Todo{}, "todo_list_id=?", this.Id)

	var mediaAttachment MediaAttachment
	db.Model(this).Related(&mediaAttachment, "MediaAttachment")
	if mediaAttachment.Id != 0 {
		db.Delete(&mediaAttachment)
	}

	return
}
