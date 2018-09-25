package model

import (
	"encoding/gob"
	"golang-todolist/frame"
)

func init() {
	// allow model.User struct to be stored in session
	gob.Register(&User{})
}

func GetUserFromSession() *User {
	val := frame.SessionGetVar("user")
	var user = &User{}
	var ok bool
	if user, ok = val.(*User); !ok {
		return nil
	}
	return user
}

// implements frame.Record
type User struct {
	Id uint `sql:"type:int PRIMARY KEY"`
	Username string `sql:"type:varchar(100)"`
	Email string `sql:"type:varchar(100)"`
	PasswordHash string `sql:"type:char(60)"`
}
