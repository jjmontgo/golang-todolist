package model

import (
	"golang-todolist/frame"
	"golang-todolist/frame/aws"
)

type TodoList struct {
	Id uint `sql:"type:int PRIMARY KEY" json:"id"`
	Name string `sql:"type:varchar(250)" json:"name"`
	Todos []Todo `json:"todos"`
	ImgSrc string `sql:"-" json:"img_src"` // populated by TodoList.GetImgSrc()
	MediaAttachment *MediaAttachment `gorm:"polymorphic:Ref;" json:"media_attachment"`
}

func (this *TodoList) GetImgSrc() string {
	if this.MediaAttachment == nil {
		return ""
	}

	return aws.SignS3ObjectUrl(this.MediaAttachment.AwsS3ObjectKey)
}

func (this *TodoList) AfterFind() (err error) {
	this.ImgSrc = this.GetImgSrc()
	return
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
