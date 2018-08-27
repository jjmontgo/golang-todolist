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
			user = &model.User{Id: "", Username: "", Email: ""}
		}
		this.Render("user/edit", "User", user)
	}

	this.Actions["Save"] = func() {
		user := model.User{Id: this.Param("id"), Username: this.Param("username"), Email: this.Param("email"),}
		// todo: validation
		err := frame.SaveRecord(&user)
		if err != nil {
			this.Error(err.Error())
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
