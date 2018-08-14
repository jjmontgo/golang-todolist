package frame

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Controller struct {
	Name string
	Actions map[string]func()
	Response http.ResponseWriter
	Request *http.Request
	Router *mux.Router
}

func NewController(name string) *Controller {
	newController := &Controller{Name: name}
	newController.Actions = make(map[string]func())
	ControllerMap[name] = newController
	return newController
}

func (this *Controller) Render(templateName string, vars interface{}) {
	view := ViewMgr.Get(templateName)
	view.Router = this.Router
	view.Render(this.Response, vars)
}
