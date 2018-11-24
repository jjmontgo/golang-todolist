package model

import "encoding/gob"

func init() {
	// allow model.User struct to be stored in session
	gob.Register(&User{})
}

// implements frame.Record
type User struct {
	Id uint `sql:"type:int PRIMARY KEY" json:"id"`
	Username string `sql:"type:varchar(100)" json:"username"`
	Email string `sql:"type:varchar(100)" json:"email"`
	PasswordHash string `sql:"type:char(60)" json:"password_hash"`
}
