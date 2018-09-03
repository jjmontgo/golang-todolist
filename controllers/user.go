package controllers

import (
	"golang-todolist/frame"
	"golang-todolist/model"
)

func init() {
	this := frame.NewController("User")

	this.Actions["Index"] = func() {
		users := model.FindUsers()
		this.Render("user/index", "Users", users)
	}

	this.Actions["Edit"] = func() {
		id := this.Param("id")
		var user *model.User
		if id != "" {
			user = model.FindUser("id", id)
		}
		if user == nil {
			user = &model.User{Username: "", Email: ""}
		}
		this.Render("user/edit", "User", user)
	}

	this.Actions["Save"] = func() {
		var user *model.User
		id := this.Param("id")
		if id == "0" {
			user = &model.User{
				Username: this.Param("username"),
				Email: this.Param("email")}
		} else {
			user = model.FindUser("id", id)
			user.Username = this.Param("username")
			user.Email = this.Param("email")
		}
		// todo: validation
		if this.Param("password") != "" {
			user.PasswordHash = frame.HashPassword(this.Param("password"))
		}
		err := frame.SaveRecord(user)
		if err != nil {
			this.Error(err)
			return
		}
		this.Redirect(frame.URL("users"))
	}

	this.Actions["Delete"] = func() {
		id := this.Param("id")
		user := model.FindUser("id", id)
		frame.DeleteRecord(user)
		this.Redirect(frame.URL("users"))
	}
}
