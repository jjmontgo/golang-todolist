package model

import (
	"log"
	"upper.io/db.v2" // required for db.Collection
	"golang-todolist/frame"
)

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
		log.Fatalf("rs.One(&user): %q\n", err)
	}
	return user
}

func Users() db.Collection {
	return frame.DB().Collection("user")
}

// implements frame.Record
type User struct {
	Id string `db:"id"`
	Username string `db:"username"`
	Email string `db:"email"`
	PasswordHash string `db:"password_hash"`
}

// frame.Record interface
func (this *User) PrimaryKey() string {
	return this.Id
}

// frame.Record interface
func (this *User) SetPrimaryKey(id string) {
	this.Id = id
}

// frame.Record interface
func (this *User) Collection() db.Collection {
	return Users()
}

