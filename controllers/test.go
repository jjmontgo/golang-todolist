package controllers

import "golang-todolist/frame"

func init() {
	newController := frame.NewController("TestController")
	newController.Actions["Test"] = func() {
		newController.Render("test", nil)
	}
}
