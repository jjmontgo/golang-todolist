package frame

import (
	"encoding/json"
	"html/template"
	"net/http"
	"bytes"
	"log"
	// "fmt"
)

func NewView(view *View) {
	tpl := template.New(view.Name)
	tpl.Funcs(template.FuncMap{
		"to_string": ToString, // frame/helpers.go
		"uint_to_string": UintToString,
		"json_encode": func(v interface {}) template.HTML {
		  a, _ := json.Marshal(v)
		  return template.HTML(a)
		}})
	if view.LayoutTemplateName == "" {
		view.LayoutTemplateName = "layout"
	}
	tpl, _ = tpl.Parse(view.Template)
	view.ParsedTemplate = tpl
	Registry.Views[view.Name] = view
}

type View struct {
	Name string
	Template string
	HasLayout bool
	LayoutTemplateName string // default to "layout"
	ParsedTemplate *template.Template

	// must be set by controller
	Response http.ResponseWriter
}

func (this *View) Render(params ...interface{}) {
	vars := BuildParameterMap(params...) // frame/helpers.go

	if this.HasLayout {
		// buffer the template to renderedContent
		var renderedContent bytes.Buffer
		if err := this.ParsedTemplate.Execute(&renderedContent, vars); err != nil {
			log.Fatalf("this.ParsedTemplate.Execute(): %q\n", err)
		}
		// render the content to the layout as html
		layoutView := Registry.Views[this.LayoutTemplateName]
		html := template.HTML(renderedContent.String())
		layoutView.Render("Content", html)
	} else {
		// render the template to the response with no layout
		this.ParsedTemplate.Execute(this.Response, vars)
	}
}
