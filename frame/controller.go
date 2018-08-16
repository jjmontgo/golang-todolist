package frame

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
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

func (this *Controller) Render(templateName string, vars interface{}) {
	view := Registry.Views[templateName]
	view.Render(vars)
}

// remove duplication with frame.view
func (this *Controller) Route(name string) string {
	url, err := Registry.Router.Get(name).URL()
	if (err != nil) {
		log.Fatalf("Registry.Router.Get(name).URL(): %q\n", err)
	}
	return url.String()
}

func (this *Controller) Redirect(url string) {
	http.Redirect(Registry.Response, Registry.Request, url, 301)
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
