
package frame

import (
	"log"
	"net/http"
	"github.com/gorilla/mux" // required by controller.Param()
	"github.com/jinzhu/gorm" // required by controller.DB()
)

type Controller struct {
	Name string
	Actions map[string]func()
	IsAccessible func(string) bool
}

func NewController(name string) *Controller {
	if c, ok := Registry.Controllers[name]; ok {
		panic("NewController(): The controller named '" + c.Name + "' already exists");
	}

	newController := &Controller{Name: name}
	newController.Actions = make(map[string]func())
	// full access by default
	newController.IsAccessible = func(actionName string) bool { return true }

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

/**
 * Typically used for retrieving an id from the request as a uint
 */
func (this *Controller) ParamUint(name string) uint {
	strParam := this.Param(name)
	return StringToUint(strParam) // see frame/helpers.go
}

func (this *Controller) Error(error error) {
	http.Error(Registry.Response, error.Error(), 500)
	log.Fatal(error)
}

func (this *Controller) DB() *gorm.DB {
	return GORM() // see frame/gorm.go
}

func (this *Controller) Email(to string, subject string, body string, from string) {
	Email(to, subject, body, from)
}
