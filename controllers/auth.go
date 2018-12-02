package controllers

import (
	"golang-todolist/frame"
	"golang-todolist/model"
)

func init() {
	this := frame.NewController("Auth")

	this.Actions["ValidateLogin"] = func() {
		isError := false

		if this.Param("username") == "" {
			isError = true
		}

		if !isError {
			var user model.User
			this.DB().Where("username = ?", this.Param("username")).First(&user)
			if user.Id == 0 {
				isError = true
			} else {
				isValidPassword := frame.VerifyPassword(this.Param("password"), user.PasswordHash)
				if !isValidPassword {
					isError = true
				} else {
					this.SessionSetVar("user", user)
					this.RenderJSON("message", "Login succeeded")
				}
			}
		}

		if isError {
			this.RenderJsonError("message", "Invalid username or password")
		}
	}

	this.Actions["Logout"] = func() {
		this.SessionSetVar("user", nil)
		this.RenderJSON("message", "You have successfully logged out.")
	}
}
