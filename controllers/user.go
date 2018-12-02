package controllers

import (
	"golang-todolist/frame"
	"golang-todolist/model"
)

func init() {
	this := frame.NewController("User")

	this.Actions["Index"] = func() {
		var users []model.User
		this.DB().Find(&users)
		this.RenderJSON("users", users)
	}

	this.Actions["Edit"] = func() {
		id := this.ParamUint("id")
		var user model.User
		if id == 0 {
			// new user
			user = model.User{Username: "", Email: ""}
		} else {
			this.DB().First(&user, id)
		}
		this.RenderJSON("user", user)
	}

	this.Actions["Save"] = func() {
		var user model.User
		id := this.ParamUint("id")
		if id == 0 {
			user = model.User{
				Username: this.Param("username"),
				Email: this.Param("email")}
		} else {
			this.DB().First(&user, id)
			user.Username = this.Param("username")
			user.Email = this.Param("email")
		}
		// todo: validation
		if this.Param("password") != "" {
			user.PasswordHash = frame.HashPassword(this.Param("password"))
		}
		this.DB().Save(&user)
	}

	this.Actions["Delete"] = func() {
		id := frame.StringToUint(this.Param("id"))
		user := model.User{Id: id}
		this.DB().Delete(&user)
	}
}
