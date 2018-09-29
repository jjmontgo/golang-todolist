package model

import (
	"time"
	"golang-todolist/frame/aws" // delete s3 object
)

type MediaAttachment struct {
	Id uint `sql:"type:int PRIMARY KEY"`
	AwsS3ObjectKey string `sql:"type:char(150)"`
	Category string `sql:"type:char(50)"`
	RefId uint `sql:"type:int"`
	RefType string `sql:"type:char(50)"`
	CreatedAt time.Time	`sql:"type:datetime"`
}

// delete the s3 object as well
func (this *MediaAttachment) BeforeDelete() (err error) {
	aws.DeleteS3Object(this.AwsS3ObjectKey)
	return
}
