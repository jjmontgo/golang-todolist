package controllers

import (
	"golang-todolist/frame"
	"golang-todolist/model"
)

func init() {
	this := frame.NewController("Auth")

	this.Actions["Login"] = func() {
		if frame.UserIsLoggedIn() {
			this.Redirect(frame.URL("index"))
			return
		}

		this.Render("auth/login")
	}

	this.Actions["ValidateLogin"] = func() {
		isError := false
		errorMessage := ""

		if this.Param("username") == "" {
			isError = true
		}

		if !isError {
			user := model.FindUser("username", this.Param("username"))
			if user == nil {
				isError = true
			} else {
				isValidPassword := frame.VerifyPassword(this.Param("password"), user.PasswordHash)
				if !isValidPassword {
					isError = true
				} else {
					frame.SessionSetVar("user", user)
					this.Redirect(frame.URL("index"))
				}
			}
		}

		if isError {
			errorMessage = "Invalid username or password"
		}

		this.Render("auth/login",	"Error", errorMessage)
	}

	this.Actions["Logout"] = func() {
		frame.SessionSetVar("user", nil)
		this.Redirect(frame.URL("login"))
	}
}
