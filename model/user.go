package model

import (
	"encoding/gob"
	"log"
	"upper.io/db.v2" // required for db.Collection
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

func FindUsers(searchParams ...interface{}) []User {
	resultSet := Users().Find(searchParams...)
	var users []User
	err := resultSet.All(&users)
	if err != nil {
		log.Fatalf("FindUsers(): %q\n", err)
	}
	return users
}

func FindUser(searchParams ...interface{}) *User {
	var user *User
	rs := Users().Find(searchParams...)
	err := rs.One(&user)
	if (err != nil) {
		return nil
	}
	return user
}

func Users() db.Collection {
	return frame.DB().Collection("user")
}

// implements frame.Record
type User struct {
	Id uint `db:"id"`
	Username string `db:"username"`
	Email string `db:"email"`
	PasswordHash string `db:"password_hash"`
}

// frame.Record interface
func (this *User) PrimaryKey() uint {
	return this.Id
}

// frame.Record interface
func (this *User) SetPrimaryKey(id uint) {
	this.Id = id
}

// frame.Record interface
func (this *User) Collection() db.Collection {
	return Users()
}

