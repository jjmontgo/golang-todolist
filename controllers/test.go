package controllers

import (
	"net/http"
	"golang-todolist/frame"
)

var TestController TestControllerType

func init() {
	TestController = TestControllerType{}
}

type TestControllerType struct {
	frame.Controller
}

func (this *TestControllerType) Test(w http.ResponseWriter, r *http.Request) {
	this.Render(w, "test", nil)
}
