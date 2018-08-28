package controllers

import (
	"golang-todolist/frame"
	"golang-todolist/model"
)

func init() {
	this := frame.NewController("Auth")

	this.Actions["Login"] = func() {
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
				}
			}
		}

		if isError {
			errorMessage = "Invalid username or password"
		}

		this.Render("auth/login", "Error", errorMessage)
	}
}
