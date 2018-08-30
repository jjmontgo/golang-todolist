package frame

import (
	"net/http"
	"github.com/gorilla/mux" // required by controller.Param()
	// "fmt"
)

type Controller struct {
	Name string
	Actions map[string]func()
}

func NewController(name string) *Controller {
	newController := &Controller{Name: name}
	newController.Actions = make(map[string]func())
	Registry.Controllers[name] = newController
	return newController
}

func (this *Controller) Render(templateName string, params ...interface{}) {
	view := Registry.Views[templateName]
	view.Render(params...)
}

func (this *Controller) Redirect(url string) {
	http.Redirect(Registry.Response, Registry.Request, url, 302)
}

func (this *Controller) Param(name string) string {
	// try to get it from GET or POST variables
	param := Registry.Request.FormValue(name)
	if param == "" {
		// try to get it from the URL
		vars := mux.Vars(Registry.Request)
		param = vars[name]
	}
	return param
}

func (this *Controller) Error(errorMessage string) {
	http.Error(Registry.Response, errorMessage, 500)
}
