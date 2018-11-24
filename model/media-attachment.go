package model

import (
	"time"
	"golang-todolist/frame/aws" // delete s3 object
)

type MediaAttachment struct {
	Id uint `sql:"type:int PRIMARY KEY" json:"id"`
	AwsS3ObjectKey string `sql:"type:char(150)" json:"aws_s3_object_key"`
	Category string `sql:"type:char(50) json:"category"`
	RefId uint `sql:"type:int" json:"ref_id"`
	RefType string `sql:"type:char(50)" json:"ref_type"`
	CreatedAt time.Time	`sql:"type:datetime" json:"created_at"`
}

// delete the s3 object as well
func (this *MediaAttachment) BeforeDelete() (err error) {
	aws.DeleteS3Object(this.AwsS3ObjectKey)
	return
}
