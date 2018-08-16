package frame

import (
	"html/template"
	"bytes"
	"log"
	// "fmt"
)

func NewView(view *View) {
	tpl := template.New(view.Name)
	tpl.Funcs(template.FuncMap{
		"route": view.Route,
		"tostring": ToString,
	})
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
}

func (this *View) Render(vars interface{}) {
	if this.HasLayout {
		// buffer the template to renderedContent
		var renderedContent bytes.Buffer
		if err := this.ParsedTemplate.Execute(&renderedContent, vars); err != nil {
			log.Fatalf("this.ParsedTemplate.Execute(): %q\n", err)
		}
		// render the content to the layout as html
		layoutView := Registry.Views[this.LayoutTemplateName]
		html := template.HTML(renderedContent.String())
		layoutView.Render(html)
	} else {
		// render the template to the response with no layout
		this.ParsedTemplate.Execute(Registry.Response, vars)
	}
}

// remove duplication with controller
func (this *View) Route(name string, vars ...string) string {
	url, err := Registry.Router.Get(name).URL(vars...)
	if (err != nil) {
		log.Fatalf("Registry.Router.Get(name).URL(): %q\n", err)
	}
	return url.String()
}
