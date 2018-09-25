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

// TODO: add this to beforeDelete gorm hook
// delete the s3 object as well
func (this *MediaAttachment) Delete() {
	aws.DeleteS3Object(this.AwsS3ObjectKey)
}
