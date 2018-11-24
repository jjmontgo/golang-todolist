
package frame

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux" // required by controller.Param()
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm" // required by controller.DB()
)

type Controller struct {
	// Name, Actions, IsAccessible are initialized at app startup time in
	// controllers/controllername.go init()
	//
	Name string
	Actions map[string]func()
	IsAccessible func(string) bool

	// Request and Response are initialized at dispatch time in dispatch.go
	// and set on a shallow copy of the controller
	Request *http.Request
	Response http.ResponseWriter
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

func (this *Controller) AfterAction() {
	this.SessionSave() // saves session data if any was added; update expiry
}

func (this *Controller) Render(templateName string, params ...interface{}) {
	// otherwise render the template
	view := Registry.Views[templateName]
	view.Render(this.Response, params...)
}

func (this *Controller) RenderJSON(params ...interface{}) {
	this.Response.Header().Set("Content-Type", "application/json")
	parameterMap := BuildParameterMap(params...)
	jsonString, err := json.Marshal(parameterMap)
	if err != nil {
		log.Fatalf("Failed to marshal json parameter map: %v", err)
	}
	this.SessionSave() // Allows the session cooke to be added to headers
	this.Response.Write([]byte(jsonString))
}

func (this *Controller) RenderJsonError(params ...interface{}) {
	// why do I have to set Content-Type before Status?
	this.Response.Header().Set("Content-Type", "application/json")
	this.HttpStatus(http.StatusUnauthorized)
	this.RenderJSON(params...)
}

func (this *Controller) Redirect(url string) {
	this.AfterAction() // action is ended early

	http.Redirect(this.Response, this.Request, url, 302)
}

func (this *Controller) HttpStatus(status int) {
	this.Response.WriteHeader(status)
}

func (this *Controller) Param(name string) string {
	// try to get it from GET or POST variables
	param := this.Request.FormValue(name)
	if param == "" {
		// try to get it from the URL
		vars := mux.Vars(this.Request)
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
	http.Error(this.Response, error.Error(), 500)
	log.Fatal(error)
}

func (this *Controller) DB() *gorm.DB {
	return GORM() // see frame/gorm.go
}

func (this *Controller) Email(to string, subject string, body string, from string) {
	Email(to, subject, body, from)
}

func (this *Controller) URL(name string, vars ...string) string {
	url, err := Registry.Router.Get(name).URL(vars...)
	if (err != nil) {
		log.Fatalf("Registry.Router.Get(name).URL(): %q\n", err)
	}
	apiGatewayPathPrefix := os.Getenv("API_GATEWAY_PATH_PREFIX")
	return apiGatewayPathPrefix + url.String()
}

// http:// or https://
func (this *Controller) Scheme() string {
	var scheme string
	isUsingTLS := this.Request.TLS != nil
	isProxiedHttps := this.Request.Header.Get("X-Forwarded-Proto") == "https"
	if isUsingTLS || isProxiedHttps {
		scheme = "https://"
	} else {
		scheme = "http://"
	}
	return scheme
}

func (this *Controller) AbsoluteURL(name string, vars ...string) string {
	scheme := this.Scheme()
	relativeURL := this.URL(name, vars...)
	return scheme + this.Request.Host + relativeURL
}

func (this *Controller) SessionSetVar(field string, value interface{}) {
	session := this.GetSession()
	session.Values[field] = value
}

func (this *Controller) SessionGetVar(field string) interface{} {
	session := this.GetSession()
	return session.Values[field]
}

func (this *Controller) GetSession() *sessions.Session {
	session, err := GetSessionStore().Get(this.Request, os.Getenv("SESSION_NAME"))
	if err != nil {
		log.Fatalf("GetSession(): %q\n", err)
	}
	return session
}

// called by dispatch.go
func (this *Controller) SessionSave() {
	session := this.GetSession()
	err := session.Save(this.Request, this.Response)
	if err != nil {
		log.Fatalf("SessionSave(): %q\n", err)
	}
}

func (this *Controller) UserIsLoggedIn() bool {
	return this.SessionGetVar("user") != nil
}
