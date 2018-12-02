package controllers

import (
	"os"
	"github.com/gorilla/mux"
	"golang-todolist/frame"
)

func init() {
	this := frame.NewController("App")

	// Make all backend routes available in JSON
	this.Actions["Urls"] = func() {
		routes := make(map[string]string)
		frame.Registry.Router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			path, _ := route.GetPathTemplate()
			name := route.GetName()
			prefix := os.Getenv("API_GATEWAY_PATH_PREFIX")
			routes[name] = this.Scheme() + this.Request.Host + prefix + path
			return nil
		})
		this.Response.Header().Set("Content-Type", "text/javascript")
		this.Render("app/urls", "Routes", routes)
	}

	// Make all translations available in JSON
	this.Actions["Translations"] = func() {
		lang := "en"
		this.Response.Header().Set("Content-Type", "text/javascript")
		this.Render("app/translations", "Translations", frame.Registry.Translations[lang])
	}
}
