package frame

import "net/http"

type Controller struct {
	Name string
	Actions map[string]func()
	Response http.ResponseWriter
	Request *http.Request
}

func NewController(name string) *Controller {
	newController := &Controller{Name: name}
	newController.Actions = make(map[string]func())
	ControllerMap[name] = newController
	return newController
}

func (this *Controller) Render(templateName string, vars interface{}) {
	ViewMgr.Get(templateName).Render(this.Response, vars)
}
